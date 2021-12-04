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

// calculateNewHorizontalPosition calculates the new horizontal position.
func calculateNewHorizontalPosition(horizontalPosition int, command string, displacement int) int {
	if command == "forward" {
		return horizontalPosition + displacement
	}

	return horizontalPosition
}

// calculateNewDepth calculates the new depth.
func calculateNewDepth(depth int, command string, displacement int) int {
	if command == "down" {
		return depth + displacement
	}

	if command == "up" {
		return depth - displacement
	}

	return depth
}

// getFinalPositionAndDepth returns the final position and depth of the
// submarine.
func getFinalPositionAndDepth(txtlines []string) (int, int) {
	// set the starting horizontal position and depth
	horizontalPosition := 0
	depth := 0

	// calculate the final position and depth
	for step, eachline := range txtlines {
		// get the command and displacement
		command, displacement := getCommandAndDisplacement(eachline)

		// calculate the new horizontal position and depth
		horizontalPosition = calculateNewHorizontalPosition(horizontalPosition, command, displacement)
		depth = calculateNewDepth(depth, command, displacement)

		// print the current step, horizontal position, and depth
		fmt.Printf("[step %d]\thorizontal position: %d, depth: %d\n", step+1, horizontalPosition, depth)
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

	// get the final position and depth
	finalPosition, finalDepth := getFinalPositionAndDepth(txtlines)

	// multiply the final position by the final depth
	finalPositionAndDepth := finalPosition * finalDepth

	// print the final position, depth, and their product
	fmt.Printf("final position: %d, final depth: %d, final position x depth: %d\n", finalPosition, finalDepth, finalPositionAndDepth)
}
