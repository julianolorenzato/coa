package cpu

type Memory interface {
	Read(Address [4]byte) [4]byte
	Write(Address [4]byte, Data [4]byte)
}

type MIPS32Memory struct {
	Cache [2][256]byte
	Main  [1024]byte
}

func (mem *MIPS32Memory) Read(Address [4]byte) [4]byte {

}
