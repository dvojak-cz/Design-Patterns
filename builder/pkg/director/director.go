package director

import (
	"builder/internal/computer"
	"builder/pkg/builder"
)

type Director struct {
	builder builder.IBuilder
}

func NewDirector(builder builder.IBuilder) *Director {
	return &Director{
		builder: builder,
	}
}

func (d *Director) SetBuilder(builder builder.IBuilder) {
	d.builder = builder
}

func (d *Director) BuildComputer() *computer.Computer{
	d.builder.Reset()
	d.builder.SetCPU(4, 2.6)
	d.builder.SetCPU(8, 1.3)
	d.builder.SetMemory(16)
	d.builder.SetMemory(32)
	d.builder.SetSSD(128)
	d.builder.SetHDD(1024, 2)
	return d.builder.GetComputer()
}
func (d *Director) BuildSuperComputer() *computer.Computer{
	d.builder.Reset()
	d.builder.SetCPU(64, 2.6)
	d.builder.SetCPU(32, 1.3)
	d.builder.SetCPU(32, 1.3)
	d.builder.SetCPU(64, 2.6)
	d.builder.SetMemory(128)
	d.builder.SetMemory(128)
	d.builder.SetMemory(128)
	d.builder.SetMemory(128)
	d.builder.SetSSD(4028)
	d.builder.SetSSD(4028)
	d.builder.SetSSD(4028)
	d.builder.SetSSD(4028)
	d.builder.SetSSD(4028)
	return d.builder.GetComputer()
}