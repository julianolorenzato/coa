package memory

import "encoding/binary"

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
