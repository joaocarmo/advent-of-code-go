package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// Represents the separator between the operation and the argument.
const OPERATION_ARGUMENT_SEPARATOR = " "
// Represents a lit pixel.
const PIXEL_LIT = "#"
// Represents a dark pixel.
const PIXEL_DARK = "."
// Represents the width of the register.
const REGISTER_WIDTH = 3

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

// CRT is a struct that represents a CRT.
type CRT struct {
	Columns int
	Rows    int
}

// drawSprite draws a sprite.
func (c *CRT) drawSprite(cpu CPU) string {
	var sprite string

	for i := 0; i < c.Rows; i++ {
		for j := 0; j < c.Columns; j++ {
			cycle := i * c.Columns + j
			register := cpu.history[cycle]

			if helpers.AbsInt(register - j) < REGISTER_WIDTH - 1 {
				sprite += PIXEL_LIT
			} else {
				sprite += PIXEL_DARK
			}
		}
		sprite += "\n"
	}

	return sprite
}

// newCRT returns a new CRT.
func newCRT(columns int, rows int) CRT {
	return CRT{
		Columns: columns,
		Rows: rows,
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
	history := cpu.getHistoryAt([]int{20, 60, 100, 140, 180, 220})
	sumSignalStrength := calcSumSignalStrengths(history)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		sumSignalStrength,
	)

	// part 2
	crt := newCRT(40, 6)
	sprite := crt.drawSprite(cpu)
	fmt.Printf(
		"[Part Two] The answer is:\n%s\n",
		sprite,
	)
}
