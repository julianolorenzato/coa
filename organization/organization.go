package organization

import (
	"github.com/julianolorenzato/fibit/cpu"
	"github.com/julianolorenzato/fibit/memory"
)

type Organization struct {
	cpu.CPU
	memory.Memory
}

func NewOrganization(cpu cpu.CPU, memory memory.Memory) *Organization {
	return &Organization{
		CPU:    cpu,
		Memory: memory,
	}
}
