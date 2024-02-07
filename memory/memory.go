package memory

import "encoding/binary"

type Word [4]byte

type Memory interface {
	Read(address any) Word
	Write(address any, word Word)
}

// RAM has 32bits address and 4GB data storage
type RAM struct {
	data [4 * GIGA]byte
}

type RAMAddress [4]byte

// Cache has 16 bits address and 256KB data storage
type Cache struct {
	data     [64 * KILO][4]byte // 32bits size
	tag      [64 * KILO][3]byte // 19bits size (32 - 11 - 2)
	validity [64 * KILO]bool    // 1bit size
}

type CacheAddress [2]byte

func (c *Cache) verifyValidity(address CacheAddress) bool {
	addr := binary.LittleEndian.Uint64(address[0:2])
	return c.validity[addr]
}

func (c *Cache) verifyTag(address CacheAddress) bool {
	addr := binary.LittleEndian.Uint64(address[0:2])
	return c.tag[addr] ==
}

type MIPS32Memory struct {
	Cache
	RAM
}

func (m *MIPS32Memory) Read(address RAMAddress) Word {
	decoded := binary.LittleEndian.Uint64(address[0:4])

	cachePos := decoded % 64 * KILO

	cacheAddress := make([]byte, cachePos)

	binary.LittleEndian.PutUint64(cacheAddress, cachePos)

	cacheAddressInBytes := [2]byte{cacheAddress[0], cacheAddress[1]}

	if m.Cache.verifyValidity(cacheAddressInBytes) && m.Cache.verifyTag(cacheAddressInBytes) {
		return m.Cache.data[cachePos]
	} else {

	}
}
