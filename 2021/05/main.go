package main

import (
	"fmt"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

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
