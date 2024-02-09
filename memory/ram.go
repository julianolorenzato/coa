package memory

// RAM has 32bits address and 4GB data storage
type RAM struct {
	data []byte
	PageTable
}

//func NewRAM(size Size) {
//	sizeLen := len(size)
//
//	math.Log2(float64(sizeLen))
//}
