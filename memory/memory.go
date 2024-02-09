package memory

type Word [4]byte

//type Memory interface {
//	Read(address any) Word
//	Write(address any, word Word)
//}

type MemorySystem struct {
	*Cache
	*TLB
	*RAM
}

type Memory[S Space] struct {
	Data S
}

func NewMemory[S Space]() *Memory[S] {
	return new(Memory[S])
}

func receiveInput[S Space](size Size) *Memory[S] {
	switch size {
	case S1KILO:
		return NewMemory()
	}
}
