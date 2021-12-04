package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// readFile reads the file and returns a slice of strings.
func readFile(filename string) []string {
	// open the file
	file, err := os.Open(filename)

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	// make a scanner to read from the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	// read each line of the file
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	return txtlines
}

// getCommandAndDisplacement returns the command and displacement from a well
// formed string.
func getCommandAndDisplacement(line string) (string, int) {
	// get the command
	command := ""
	displacement := 0

	// split the string using the space as the delimiter
	split := strings.Split(line, " ")

	// get the command
	command = split[0]

	// get the displacement
	displacement, _ = strconv.Atoi(split[1])

	return command, displacement
}

// calculateNewAim calculates the new aim.
func calculateNewAim(aim int, command string, displacement int) int {
	if command == "down" {
		return aim + displacement
	}

	if command == "up" {
		return aim - displacement
	}

	return aim
}

// calculateNewHorizontalPosition calculates the new horizontal position.
func calculateNewHorizontalPosition(horizontalPosition int, command string, displacement int) int {
	if command == "forward" {
		return horizontalPosition + displacement
	}

	return horizontalPosition
}

// calculateNewDepthPartOne calculates the new depth.
func calculateNewDepthPartOne(depth int, command string, displacement int) int {
	if command == "down" {
		return depth + displacement
	}

	if command == "up" {
		return depth - displacement
	}

	return depth
}

// calculateNewDepthPartTwo calculates the new depth.
func calculateNewDepthPartTwo(aim int, depth int, command string, displacement int) int {
	if command == "forward" {
		return depth + aim*displacement
	}

	return depth
}

// getFinalPositionAndDepthPartOne returns the final position and depth of the
// submarine.
func getFinalPositionAndDepthPartOne(txtlines []string) (int, int) {
	// set the starting horizontal position and depth
	horizontalPosition := 0
	depth := 0

	// calculate the final position and depth
	for step, eachline := range txtlines {
		// get the command and displacement
		command, displacement := getCommandAndDisplacement(eachline)

		// calculate the new horizontal position and depth
		horizontalPosition = calculateNewHorizontalPosition(horizontalPosition, command, displacement)
		depth = calculateNewDepthPartOne(depth, command, displacement)

		// print the current step, horizontal position, and depth
		fmt.Printf("[step %d]\thorizontal position: %d, depth: %d\n", step+1, horizontalPosition, depth)
	}

	return horizontalPosition, depth
}

// getFinalPositionAndDepthPartTwo returns the final position and depth of the
// submarine.
func getFinalPositionAndDepthPartTwo(txtlines []string) (int, int) {
	// set the starting horizontal position and depth
	aim := 0
	horizontalPosition := 0
	depth := 0

	// calculate the final position and depth
	for step, eachline := range txtlines {
		// get the command and displacement
		command, displacement := getCommandAndDisplacement(eachline)

		// calculate the new horizontal position and depth
		aim = calculateNewAim(aim, command, displacement)
		horizontalPosition = calculateNewHorizontalPosition(horizontalPosition, command, displacement)
		depth = calculateNewDepthPartTwo(aim, depth, command, displacement)

		// print the current step, aim, horizontal position, and depth
		fmt.Printf("[step %d]\taim: %d, horizontal position: %d, depth: %d\n", step+1, aim, horizontalPosition, depth)
	}

	return horizontalPosition, depth
}

func main() {
	// get the filename from the command line
	args := os.Args[1:]

	if len(args) != 1 {
		log.Fatal("you must supply a filename")
	}

	// read the file
	filename := args[0]
	txtlines := readFile(filename)

	// get the final position and depth (Part One)
	finalPosition, finalDepth := getFinalPositionAndDepthPartOne(txtlines)

	// multiply the final position by the final depth (Part One)
	finalPositionAndDepth := finalPosition * finalDepth

	// print the final position, depth, and their product (Part One)
	fmt.Printf("[Part One] final position: %d, final depth: %d, final position x depth: %d\n", finalPosition, finalDepth, finalPositionAndDepth)

	// get the final position and depth (Part Two)
	finalPosition, finalDepth = getFinalPositionAndDepthPartTwo(txtlines)

	// multiply the final position by the final depth (Part Two)
	finalPositionAndDepth = finalPosition * finalDepth

	// print the final position, depth, and their product (Part Two)
	fmt.Printf("[Part Two] final position: %d, final depth: %d, final position x depth: %d\n", finalPosition, finalDepth, finalPositionAndDepth)
}
