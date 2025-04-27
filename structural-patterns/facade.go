package main

// Refer to: https://refactoring.guru/design-patterns/facade

/*
Facade is a structural design pattern that provides
a simplified interface to a library, a framework, or any other complex set of classes.
*/

import "fmt"

// CPU Subsystem 1
type CPU struct{}

func (c *CPU) Freeze() {
	fmt.Println("Freezing processor...")
}

func (c *CPU) Jump(position int) {
	fmt.Printf("Jumping to position %d...\n", position)
}

func (c *CPU) Execute() {
	fmt.Println("Executing instructions...")
}

// Memory Subsystem 2
type Memory struct{}

func (m *Memory) Load(position int, data string) {
	fmt.Printf("Loading '%s' into memory at position %d...\n", data, position)
}

// HardDrive Subsystem 3
type HardDrive struct{}

func (hd *HardDrive) Read(position int) string {
	fmt.Printf("Reading from hard drive at position %d...\n", position)
	return "bootloader"
}

// Computer Facade
type Computer struct {
	cpu       *CPU
	memory    *Memory
	hardDrive *HardDrive
}

func NewComputer() *Computer {
	return &Computer{
		cpu:       &CPU{},
		memory:    &Memory{},
		hardDrive: &HardDrive{},
	}
}

func (c *Computer) Start() {
	fmt.Println("Starting computer using Facade...")
	c.cpu.Freeze()
	data := c.hardDrive.Read(0)
	c.memory.Load(0, data)
	c.cpu.Jump(0)
	c.cpu.Execute()
}

// Client
func main() {
	computer := NewComputer()
	computer.Start()
}
