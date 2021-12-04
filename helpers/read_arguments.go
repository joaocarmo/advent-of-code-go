package helpers

import (
	"os"
)

// ReadArguments reads the arguments from the command line.
func ReadArguments() []string {
	// get the filename from the command line
	args := os.Args[1:]

	if len(args) != 1 {
		os.Stderr.WriteString("you must supply a filename\n")
		os.Exit(1)
	}

	return args
}
