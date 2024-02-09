package cpu

type CPU interface {
	ExecuteInstr(instr [4]byte)
}
