package helpers

import (
	"bufio"
	"fmt"
	"os"
)

// ReadFile reads the file and returns a slice of strings.
func ReadFile(filename string) []string {
	// open the file
	file, err := os.Open(filename)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed opening file: %s\n", err)
		os.Exit(1)
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
