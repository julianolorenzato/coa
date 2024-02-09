package mips32

import (
	"encoding/binary"
	"errors"
	"github.com/julianolorenzato/fibit/memory"
	"io"
	"os"
)

// Bitmasks (&)
const (
	OPCODE   = 0b_111111_00000_00000_00000_00000_000000
	RS       = 0b_000000_11111_00000_00000_00000_000000
	RT       = 0b_000000_00000_11111_00000_00000_000000
	RD       = 0b_000000_00000_00000_11111_00000_000000
	SHAMT    = 0b_000000_00000_00000_00000_11111_000000
	FUNCTION = 0b_000000_00000_00000_00000_00000_111111
	OFFSET   = 0b_000000_00000_00000_11111_11111_111111
	TARGET   = 0b_000000_11111_11111_11111_11111_111111
	HI       = 0b_111111_11111_11111_00000_00000_000000
	LO       = 0b_000000_00000_00000_11111_11111_111111
)

type MIPS32 struct {
	Registers    [32]uint32
	Hi           uint32
	Lo           uint32
	Pc           uint32
	InstrMemory  memory.Memory
	DataMemory   memory.Memory
	Running      bool
	Instructions map[uint32]func(info *instrInfo)
}

func NewMIPS32() *MIPS32 {
	mips32 := &MIPS32{
		Registers:    [32]uint32{},
		Hi:           0,
		Lo:           0,
		Pc:           0,
		InstrMemory:  nil,
		DataMemory:   nil,
		Running:      false,
		Instructions: map[uint32]func(info *instrInfo){},
	}

	mips32.fillInstructions()

	return mips32
}

func (cpu *MIPS32) fillInstructions() {
	cpu.Instructions = map[uint32]func(info *instrInfo){
		0b_000000: cpu.r,
		0b_000100: cpu.beq,
		0b_000010: cpu.j,
		0b_111111: cpu.hlt,
	}
}

func (cpu *MIPS32) LoadInstructionsInMemory(filename string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 4)
	var addressCounter uint32 = 0
	for {
		n, err := fd.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		cpu.InstrMemory.Write(encodeWord(addressCounter), [4]byte(b))
		addressCounter += 4

		if n == 0 {
			break
		}
	}
}

func (cpu *MIPS32) StartProgram() {
	cpu.Running = true

	for cpu.Running {
		instr := cpu.InstrMemory.Read(encodeWord(cpu.Pc))
		cpu.executeInstr(instr)
		cpu.Pc += 4
	}
}

func encodeWord(data uint32) [4]byte {
	buf := make([]byte, 4)
	binary.LittleEndian.PutUint32(buf, data)
	return [4]byte{buf[0], buf[1], buf[2], buf[3]}
}

func decodeWord(word [4]byte) uint32 {
	return binary.LittleEndian.Uint32(word[0:4])
}

type instrInfo struct {
	rs     uint32
	rt     uint32
	rd     uint32
	shamt  uint32
	funct  uint32
	offset uint32
	target uint32
}

func decodeInstr(data uint32) (opCode uint32, info instrInfo) {
	opCode = (data & OPCODE) >> 26
	info = instrInfo{
		rs:     (data & RS) >> 21,
		rt:     (data & RT) >> 16,
		rd:     (data & RD) >> 11,
		shamt:  (data & SHAMT) >> 6,
		funct:  data & FUNCTION,
		offset: data & OFFSET,
		target: data & TARGET,
	}
	return
}

func (cpu *MIPS32) executeInstr(instr [4]byte) {
	opCode, info := decodeInstr(decodeWord(instr))
	cpu.Instructions[opCode](&info)
}

func (cpu *MIPS32) r(info *instrInfo) {
	rOperations := map[uint32]func(info *instrInfo){
		0b_100000: cpu.add,
		0b_100100: cpu.and,
		0b_011010: cpu.div,
		0b_011000: cpu.mult,
		0b_010010: cpu.mflo,
		0b_010000: cpu.mfhi,
		0b_100111: cpu.nor,
		0b_100101: cpu.or,
		0b_101010: cpu.slt,
		0b_100010: cpu.sub,
		0b_100110: cpu.xor,
	}

	rOperations[info.funct](info)
}

func (cpu *MIPS32) add(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Registers[info.rs] + cpu.Registers[info.rt]
}

func (cpu *MIPS32) and(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Registers[info.rs] & cpu.Registers[info.rt]
}

func (cpu *MIPS32) div(info *instrInfo) {
	cpu.Lo = cpu.Registers[info.rs] / cpu.Registers[info.rt]
	cpu.Hi = cpu.Registers[info.rs] % cpu.Registers[info.rt]
}

func (cpu *MIPS32) mult(info *instrInfo) {
	res := uint64(cpu.Registers[info.rs]) * uint64(cpu.Registers[info.rt])
	cpu.Hi = uint32((res & HI) >> 32)
	cpu.Lo = uint32(res & LO)
}

func (cpu *MIPS32) mflo(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Lo
}

func (cpu *MIPS32) mfhi(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Hi
}

func (cpu *MIPS32) nor(info *instrInfo) {
	cpu.Registers[info.rd] = ^(cpu.Registers[info.rs] | cpu.Registers[info.rt])
}

func (cpu *MIPS32) or(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Registers[info.rs] | cpu.Registers[info.rt]
}

func (cpu *MIPS32) slt(info *instrInfo) {
	if cpu.Registers[info.rs] < cpu.Registers[info.rt] {
		cpu.Registers[info.rd] = 1
	} else {
		cpu.Registers[info.rd] = 0
	}
}

func (cpu *MIPS32) sub(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Registers[info.rs] + cpu.Registers[info.rt]
}

func (cpu *MIPS32) xor(info *instrInfo) {
	cpu.Registers[info.rd] = cpu.Registers[info.rs] ^ cpu.Registers[info.rt]
}

func (cpu *MIPS32) beq(info *instrInfo) {
	if cpu.Registers[info.rs] == cpu.Registers[info.rt] {
		cpu.Pc += info.offset
	}
}

func (cpu *MIPS32) j(info *instrInfo) {
	cpu.Pc = info.target
}

func (cpu *MIPS32) hlt(_ *instrInfo) {
	cpu.Running = false
}
