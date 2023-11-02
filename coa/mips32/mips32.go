package mips32

import (
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type MIPS32 struct {
	Registers   [32]uint32
	Hi          uint32
	Lo          uint32
	Pc          uint32
	InstrMemory Memory
	DataMemory  Memory
	Running     bool
}

func NewMIPS32() *MIPS32 {
	return &MIPS32{
		Registers:   [32]uint32{0},
		Hi:          0,
		Lo:          0,
		Pc:          0,
		InstrMemory: nil,
		DataMemory:  nil,
		Running:     false,
	}
}

type Instr uint32

//var instrBitmasks = map[string]uint32{
//	"add": 0b_000000_00000_00000_00000_00000_100000,
//	"and": 0b_000000_00000_00000_00000_00000_100100,
//	"nor": 0b_000000_00000_00000_00000_00000_100111,
//	"sub": 0b_000000_00000_00000_00000_00000_100010,
//}

// OPCODES
const (
	R   = 0b_000000
	BEQ = 0b_000100
	HLT = 0b_111111
)

// FUNCTS
const (
	ADD = 0b_100000
	AND = 0b_100100
	NOR = 0b_100111
	OR  = 0b_100101
	SLT = 0b_101010
	SUB = 0b_100010
	XOR = 0b_100110
)

type Memory interface {
	Read(Address uint32) int32
	Write(Address uint32, Data int32)
}

// MAYBE CONVERT ALL TYPES TO BYTE ARRAYS AND USE UINT32 OR INT32 INSIDE THE FUNCTIONS
func (cpu *MIPS32) LoadInstructionsInMemory(filename string) {
	fd, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	b := make([]byte, 4)

	var counter uint32 = 0

	for {
		read, err := fd.Read(b)
		if err != nil && !errors.Is(err, io.EOF) {
			panic(err)
		}

		instr := binary.LittleEndian.Uint32(b)
		cpu.InstrMemory.Write(counter, int32(instr))
		counter += 4

		if read == 0 {
			break
		}
	}
}

func (cpu *MIPS32) StartProgram() {
	cpu.Running = true

	for cpu.Running {
		pc := cpu.Pc
		instr := Instr(cpu.InstrMemory.Read(pc))
		cpu.DecodeInstr(instr)
		cpu.Pc += 4
	}
}

func (cpu *MIPS32) DecodeInstr(instr Instr) {
	opCode := instr >> 26
	funct := (instr << 26) >> 26

	switch opCode {
	case R:
		regSource1 := (instr << 6) >> 26
		regSource2 := (instr << 11) >> 26
		regDestiny := (instr << 16) >> 26

		switch funct {
		case ADD:
			cpu.Registers[regDestiny] = cpu.Registers[regSource1] + cpu.Registers[regSource2]
		case AND:
			cpu.Registers[regDestiny] = cpu.Registers[regSource1] & cpu.Registers[regSource2]
		case NOR:
			cpu.Registers[regDestiny] = ^(cpu.Registers[regSource1] | cpu.Registers[regSource2])
		case OR:
			cpu.Registers[regDestiny] = cpu.Registers[regSource1] | cpu.Registers[regSource2]
		case SLT:
			if cpu.Registers[regSource1] < cpu.Registers[regSource2] {
				cpu.Registers[regDestiny] = 1
			} else {
				cpu.Registers[regDestiny] = 0
			}
		case SUB:
			cpu.Registers[regDestiny] = cpu.Registers[regSource1] - cpu.Registers[regSource2]
		case XOR:
			cpu.Registers[regDestiny] = cpu.Registers[regSource1] ^ cpu.Registers[regSource2]
		}
	case BEQ:
		regSource1 := (instr << 6) >> 26
		regSource2 := (instr << 11) >> 26

		if cpu.Registers[regSource1] == cpu.Registers[regSource2] {
			cpu.Pc
		}
	case HLT:
		cpu.Running = false
	}
}

type Registers struct {
	//Zero uint32
	//At uint32
	//V0 uint32
	//V1 uint32
	//A0 uint32
	//A1 uint32
	//A2 uint32
	//A3 uint32
	//T0 uint32
	//T1 uint32
	//T2 uint32
	//T3 uint32
	//T4 uint32
	//T5 uint32
	//T6 uint32
	//T7 uint32
	//S0 uint32
	//S1 uint32
	//S2 uint32
	//S3 uint32
	//S4 uint32
	//S5 uint32
	//S6 uint32
	//S7 uint32
	//T8 uint32
	//T9 uint32
	//K0 uint32
	//K1 uint32
	//Gp uint32
	//Sp uint32
	//Fp uint32
	//Ra uint32
	Hi uint32
	Lo uint32
}
