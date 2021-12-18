package main

import (
	"fmt"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// print the text lines
	fmt.Println(txtlines)
}
