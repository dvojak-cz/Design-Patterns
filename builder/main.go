package main

import (
	"builder/pkg/builder"
	"builder/pkg/director"
	"fmt"
)

func main() {
	laptopBuilder := builder.GetBuilder(builder.LaptopBuilderType)
	desktopBuilder := builder.GetBuilder(builder.DesktopBuilderType)

	fmt.Println("I am about to build a laptop")
	director := director.NewDirector(laptopBuilder)
	laptop := director.BuildComputer()
	fmt.Println(laptop)
	superLaptop := director.BuildSuperComputer()
	fmt.Println(superLaptop)

	fmt.Println("I am about to build a desktop")
	director.SetBuilder(desktopBuilder)
	desktop := director.BuildComputer()
	fmt.Println(desktop)
	superDesktop:= director.BuildSuperComputer()
	fmt.Println(superDesktop)
}
