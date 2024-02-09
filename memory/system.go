package memory

type System struct {
	*Cache
	*TLB
	*RAM
	*Disk
}

type Info struct {
	Disk struct {
		Size uint32
	}
	RAM struct {
		Size            uint32
		BytesPerAddress uint32
	}
	Caches []struct {
		NumberOfSets      uint32
		BlockSize         uint32
		Associativity     uint32
		ReplacementPolicy rune
	}
}
