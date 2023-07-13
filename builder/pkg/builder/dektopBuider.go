package builder

import "builder/internal/computer"

type DesktopBuilder struct {
	c *computer.Computer
}

func newDesktopBuilder() DesktopBuilder {
	return DesktopBuilder{
		c: &computer.Computer{
			Type: "Desktop",
		},
	}
}
func (d DesktopBuilder) Reset() {
	d.c = &computer.Computer{
		Type: "Desktop",
	}
}
func (d DesktopBuilder) SetCPU(cores uint, fr float32) {
	d.c.CPUs = append(d.c.CPUs, computer.CPU{Cores: cores, Frequency: fr})
}
func (d DesktopBuilder) SetMemory(c uint) {
	d.c.Memory = append(d.c.Memory, computer.Memory{Capacity: c})
}
func (d DesktopBuilder) SetSSD(capacity uint) {
	d.c.Disks = append(d.c.Disks, computer.SolidStateDrive{Capacity: capacity})
}
func (d DesktopBuilder) SetHDD(capacity uint, platters uint) {
	d.c.Disks = append(d.c.Disks, computer.HardDrive{Capacity: capacity, NumberOfPlatters: platters})
}
func (d DesktopBuilder) GetComputer() *computer.Computer {
	return d.c
}
