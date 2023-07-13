package builder

import "builder/internal/computer"

type IBuilder interface {
	Reset()
	SetCPU(uint, float32)
	SetMemory(uint)
	SetSSD(uint)
	SetHDD(uint, uint)
	GetComputer() *computer.Computer
}

type ComputerBuilderType int

const (
	LaptopBuilderType ComputerBuilderType = iota
	DesktopBuilderType
)

func GetBuilder(builderType ComputerBuilderType) IBuilder {
	switch builderType {
	case LaptopBuilderType:
		return newLaptopBuilder()
	case DesktopBuilderType:
		return newDesktopBuilder()
	default:
		return nil
	}
}
