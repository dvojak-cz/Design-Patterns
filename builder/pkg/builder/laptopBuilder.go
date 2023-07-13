package builder

import "builder/internal/computer"

type LaptopBuilder struct {
	c *computer.Computer
}

func newLaptopBuilder() LaptopBuilder {
	return LaptopBuilder{
		c: &computer.Computer{
			Type: "Laptop",
		},
	}
}
func (l LaptopBuilder) Reset() {
	l.c = &computer.Computer{
		Type: "Laptop",
	}
}
func (l LaptopBuilder) SetCPU(cores uint, fr float32) {
	l.c.CPUs = []computer.CPU{
		{
			Cores:     cores,
			Frequency: fr,
		},
	}
}
func (l LaptopBuilder) SetMemory(c uint) {
	if len(l.c.Memory) < 2 {
		l.c.Memory = append(l.c.Memory, computer.Memory{Capacity: c})
	} else {
		l.c.Memory[1] = computer.Memory{Capacity: c}
	}
}
func (l LaptopBuilder) SetSSD(c uint) {
	l.c.Disks = append(l.c.Disks, computer.SolidStateDrive{Capacity: c})
}
func (l LaptopBuilder) SetHDD(_ uint, _ uint) {
}
func (l LaptopBuilder) GetComputer() *computer.Computer {
	return l.c
}
