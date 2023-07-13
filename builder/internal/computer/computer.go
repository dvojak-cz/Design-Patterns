package computer

import (
	"fmt"
	"strings"
)

type Computer struct {
	Type   string
	CPUs   []CPU
	Memory []Memory
	Disks  []Disk
}

type Disk interface {
	String() string
}

type CPU struct {
	Cores     uint
	Frequency float32
}

type Memory struct {
	Capacity uint
}

type SolidStateDrive struct {
	Capacity uint
}

type HardDrive struct {
	Capacity         uint
	NumberOfPlatters uint
}

// TODO: implement Stringer interface for all types using generics
func (c *Computer) String() string {
	var description string
	description += fmt.Sprintf("%s:\n", c.Type)

	arrayOfCPUsString := make([]string, len(c.CPUs))
	for i, cpu := range c.CPUs {
		arrayOfCPUsString[i] = cpu.String()
	}

	arrayOfMemoryStrings := make([]string, len(c.Memory))
	for i, cpu := range c.Memory {
		arrayOfMemoryStrings[i] = cpu.String()
	}
	description += fmt.Sprintf("\t CPU: [%s]\n", strings.Join(arrayOfMemoryStrings, ", "))

	arrayOfDisksStrings := make([]string, len(c.Disks))
	for i, disk := range c.Disks {
		arrayOfDisksStrings[i] = disk.String()
	}
	description += fmt.Sprintf("\t Disks: [%s]\n", strings.Join(arrayOfDisksStrings, ", "))

	return description
}
func (c *CPU) String() string {
	return fmt.Sprintf("%d crs @ %.2f GHz", c.Cores, c.Frequency)
}
func (m *Memory) String() string {
	return fmt.Sprintf("%d GB", m.Capacity)
}
func (ssd SolidStateDrive) String() string {
	return fmt.Sprintf("%d GB", ssd.Capacity)
}
func (hd HardDrive) String() string {
	return fmt.Sprintf("%d GB - %d pl", hd.Capacity, hd.NumberOfPlatters)
}
