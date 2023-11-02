package memory

type Computer[A Address, W Word] struct {
	CPU[A, W]
	MemorySystem struct {
		CacheL1 Memory[A, W]
		CacheL2 Memory[A, W]
		Main Memory[A, W]
	}
}

type InstrSet[W Word] []W

type CPU[A Address, W Word] interface {
	DecodeInstr()
}

type Memory[A Address, W Word] interface {
	Read(address A)
	Write(address A, data W)
}

type Word interface {
	[128]byte | [256]byte | [512]byte
}

type Address interface {
	[4]byte | [8]byte
}

var b = a{}

func c() {
	b.
}
