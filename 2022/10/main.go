package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// Represents the separator between the operation and the argument.
const OPERATION_ARGUMENT_SEPARATOR = " "

// Operation is an enum that represents the operation.
type Operation int

const (
	NOOP Operation = iota
	ADDX
)

// String returns the string representation of the operation.
func (o Operation) String() string {
	switch o {
	case NOOP:
		return "noop"
	case ADDX:
		return "addx"
	}
	return ""
}

// Instruction is a struct that represents an instruction.
type Instruction struct {
	operation Operation
	argument  int
}

// String returns the string representation of the instruction.
func (i Instruction) String() string {
	return fmt.Sprintf("%s %d\n", i.operation, i.argument)
}

// CPU is a struct that represents a CPU.
type CPU struct {
	cycle    int
	register int
	history  map[int]int
}

// getHistoryAt returns the history of the CPU at a given cycle.
func (c *CPU) getHistoryAt(cycles []int) [][]int {
	var history [][]int

	for _, cycle := range cycles {
		history = append(history, []int{cycle, c.history[cycle - 1]})
	}

	return history
}

// increaseCycle increases the cycle of the CPU by one.
func (c *CPU) increaseCycle() {
	c.history[c.cycle] = c.register
	c.cycle++
}

// updateRegister updates the register of the CPU.
func (c *CPU) updateRegister(argument int) {
	c.register += argument
}

// execute executes an instruction.
func (c *CPU) execute(instruction Instruction) {
	c.increaseCycle()

	switch instruction.operation {
	case NOOP:
		// `noop` takes one cycle to complete. It has no other effect.
	case ADDX:
		// `addx`` V takes two cycles to complete. After two cycles, the X register is increased by the value V.
		c.increaseCycle()
		c.updateRegister(instruction.argument)
	}
}

// executeAll executes all instructions.
func (c *CPU) executeAll(instructions []Instruction) {
	for _, instruction := range instructions {
		c.execute(instruction)
	}
}

// String returns the string representation of the CPU.
func (c CPU) String() string {
	return fmt.Sprintf("cycle: %d, register: %d\n", c.cycle, c.register)
}

// newCPU returns a new CPU.
func newCPU() CPU {
	return CPU{
		cycle:    0,
		register: 1,
		history:  make(map[int]int),
	}
}

// calcSignalStrength calculates the signal strength.
func calcSignalStrength(input []int) int {
	if len(input) == 2 {
		cycle := input[0]
		register := input[1]

		return cycle * register
	}

	return 0
}

// calcSumSignalStrengths calculates the sum of the signal strengths.
func calcSumSignalStrengths(input [][]int) int {
	sum := 0

	for _, signal := range input {
		sum += calcSignalStrength(signal)
	}

	return sum
}

// instructionFromString returns an instruction from a string.
func instructionFromString(input string) Instruction {
	var operation Operation
	var argument int
	operationArgument := strings.Split(input, OPERATION_ARGUMENT_SEPARATOR)

	switch operationArgument[0] {
	case "noop":
		operation = NOOP
	case "addx":
		operation = ADDX
	}

	if len(operationArgument) > 1 {
		argument, _ = strconv.Atoi(operationArgument[1])
	}

	return Instruction{
		operation: operation,
		argument:  argument,
	}
}

// getInstructionsFromFile returns a slice of instructions from a slice of strings.
func getInstructionsFromFile(input []string) []Instruction {
	var instructions []Instruction

	for _, line := range input {
		instructions = append(instructions, instructionFromString(line))
	}

	return instructions
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	instructions := getInstructionsFromFile(txtlines)

	// part 1
	cpu := newCPU()
	cpu.executeAll(instructions)
	fmt.Println(cpu)
	history := cpu.getHistoryAt([]int{20, 60, 100, 140, 180, 220})
	sumSignalStrength := calcSumSignalStrengths(history)
	fmt.Println(sumSignalStrength)
}
