package main

import (
	"fmt"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// absDiffInt returns the absolute difference between two integers.
func absDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// absInt returns the absolute value of the given integer.
func absInt(x int) int {
	return absDiffInt(x, 0)
}

// minOf returns the minimum of the given values.
func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

// maxOf returns the maximum of the given values.
func maxOf(vars ...int) int {
	max := vars[0]

	for _, i := range vars {
		if max < i {
			max = i
		}
	}

	return max
}

// parseFileForVents parses the file for vents.
func parseFileForVents(lines []string) []Vent {
	var vents []Vent

	for _, line := range lines {
		// parse the line
		points := strings.Split(line, " -> ")

		// create a new vent
		v := Vent{}
		v.new(points[0], points[1])

		// add the vent to the list
		vents = append(vents, v)
	}

	return vents
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// parse the file
	vents := parseFileForVents(txtlines)

	// create the board
	board := Board{}
	board.new(vents)

	// get the number of points with overlap
	overlap := board.getOverlap(2)

	// print the board
	fmt.Println(board.toString())

	// print the number of points with overlap
	fmt.Printf("The number of points with overlap is %d\n", overlap)
}
