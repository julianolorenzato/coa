package memory

type RAM struct {
	memory Memory
	//PageTable
}

func NewRAM(memory Memory) *RAM {
	return &RAM{
		memory: memory,
	}
}
