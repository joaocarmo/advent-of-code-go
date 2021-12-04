package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// convertToMatrix converts the input to a matrix of ints.
func convertToMatrix(txtlines []string) [][]int {
	// create the matrix
	matrix := make([][]int, len(txtlines))
	for i := range matrix {
		matrix[i] = make([]int, len(txtlines[i]))
	}

	// fill the matrix
	for i := range txtlines {
		for j := range txtlines[i] {
			matrix[i][j], _ = strconv.Atoi(string(txtlines[i][j]))
		}
	}

	return matrix
}

// transposeMatrix transposes a matrix.
func transposeMatrix(matrix [][]int) [][]int {
	// create the transposed matrix
	transposed := make([][]int, len(matrix[0]))
	for i := range transposed {
		transposed[i] = make([]int, len(matrix))
	}

	// fill the transposed matrix
	for i := range matrix {
		for j := range matrix[i] {
			transposed[j][i] = matrix[i][j]
		}
	}

	return transposed
}

// countBitsAndAppendBit counts the number of bits and appends the bit to the output.
func countBitsAndAppendBit(matrix [][]int, compareFn func(countZero int, countOne int) string) string {
	// create the output
	output := make([]string, len(matrix))

	// find the most common bits
	for i := range matrix {
		var countZero, countOne int
		for j := range matrix[i] {
			if matrix[i][j] == 0 {
				countZero++
			} else {
				countOne++
			}
		}

		// append the resulting bit
		output[i] = compareFn(countZero, countOne)
	}

	return strings.Join(output, "")
}

// findMostCommonBits finds the most common bits.
func findMostCommonBits(matrix [][]int) string {
	// define the compare function
	compareFn := func(countZero int, countOne int) string {
		// append the most common bit
		if countZero > countOne {
			return "0"
		} else {
			return "1"
		}
	}

	return countBitsAndAppendBit(matrix, compareFn)
}

// findLeastCommonBits finds the least common bits.
func findLeastCommonBits(matrix [][]int) string {
	// define the compare function
	compareFn := func(countZero int, countOne int) string {
		// append the least common bit
		if countZero < countOne {
			return "0"
		} else {
			return "1"
		}
	}

	return countBitsAndAppendBit(matrix, compareFn)
}

// findGammaAndEpsilon finds the gamma and epsilon rates.
func findGammaAndEpsilon(txtlines []string) (int, int) {
	var gammaBin, epsilonBin string
	var gamma, epsilon int64

	// convert the input to a matrix of ints
	matrix := convertToMatrix(txtlines)

	// transpose the matrix
	transposed := transposeMatrix(matrix)

	// find the gamma rate
	gammaBin = findMostCommonBits(transposed)
	gamma, _ = strconv.ParseInt(gammaBin, 2, 64)

	// find the epsilon rate
	epsilonBin = findLeastCommonBits(transposed)
	epsilon, _ = strconv.ParseInt(epsilonBin, 2, 64)

	return int(gamma), int(epsilon)
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// find the gamma rate and the epsilon rate
	gamma, epsilon := findGammaAndEpsilon(txtlines)

	// calculate the power consumption
	power := gamma * epsilon

	// print
	fmt.Printf("Gamma: %d\n", gamma)
	fmt.Printf("Epsilon: %d\n", epsilon)
	fmt.Printf("Power: %d\n", power)
}
