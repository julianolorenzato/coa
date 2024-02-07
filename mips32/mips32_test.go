package mips32

import (
	"testing"
)

func TestNewMIPS32(t *testing.T) {
	mips32 := NewMIPS32()

	if len(mips32.Instructions) == 0 {
		t.Errorf("Instructions not filled in mips32")
	}
}

func Test_decodeInstr(t *testing.T) {
	var instr uint32 = 0b_100110_11101_01100_11111_00010_111010
	opCode, info := decodeInstr(instr)

	var opCodeWant uint32 = 0b_100110
	infoWant := instrInfo{
		rs:     0b_11101,
		rt:     0b_01100,
		rd:     0b_11111,
		shamt:  0b_00010,
		funct:  0b_111010,
		offset: 0b_11111_00010_111010,
	}

	if opCode != opCodeWant || info != infoWant {
		t.Errorf("Not the same result, %b %b", opCode, info)
	}
}

//func TestMIPS32_executeInstr(t *testing.T) {
//	mips32 := NewMIPS32()
//	mips32.Registers[0] = 12
//	mips32.Registers[1] = 20
//	mips32.Registers[2] = 30
//
//	tests := []struct {
//		name  string
//		instr [4]byte
//		want MIPS32
//	}{
//		{"add", encodeWord(0b_000000_00000_00010_00001_00000_100000)},
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			mips32.executeInstr(test.instr)
//
//		})
//
//		mips32 = NewMIPS32()
//		mips32.Registers[0] = 12
//		mips32.Registers[1] = 20
//		mips32.Registers[2] = 30
//	}
//}
