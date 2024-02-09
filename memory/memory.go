package memory

import (
	"fmt"
)

type Memory []byte

func NewMemory(s uint32) (Memory, error) {
	size, err := validateSize(s)
	if err != nil {
		return nil, err
	}

	return make(Memory, size), nil
}

func (m Memory) Size() uint32 {
	return uint32(len(m))
}

func (m Memory) Read(address uint32) byte {
	return m[address]
}

func (m Memory) Write(address uint32, data byte) {
	m[address] = data
}

func validateSize(s uint32) (Size, error) {
	size := Size(s)

	if availableSizes[size] {
		return size, nil
	}

	return size, fmt.Errorf("invalid size %dbytes", size)
}
