package memory

type MIPS32 struct {
	Memory struct {
		Cache Memory[[256]byte]
	}
}

type Instr int32

const (
	ADD = 0b11111100000001
)

var instrSet = map[string]Instr{
	"add": 0b_1111_1111_1111_1111_11111_1111_11111,
	"sub": {0b11101100, 2, 3, 255},
	"addi"
}
type Word [byte]

type Main [65536]byte // 64KB

type CacheSize interface {
	[1024]byte | [512]byte | [216]byte
}

type Cache[T CacheSize] struct {
	Data T
}

type Cache interface {
	Read(address Address)
}

func (m *Memory) Read(addr Address) {

}

func (m *Memory) Write(addr Address) {

}
