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

// countBits counts the number of bits in a row.
func countBits(row []int) (int, int) {
	var countZero, countOne int

	for i := range row {
		if row[i] == 0 {
			countZero++
		} else {
			countOne++
		}
	}

	return countZero, countOne
}

// countBitsAndAppendBit counts the number of bits and appends the bit to the output.
func countBitsAndAppendBit(matrix [][]int, compareFn func(countZero int, countOne int) string) string {
	// create the output
	output := make([]string, len(matrix))

	// find the most common bits
	for i := range matrix {
		countZero, countOne := countBits(matrix[i])

		// append the resulting bit
		output[i] = compareFn(countZero, countOne)
	}

	return strings.Join(output, "")
}

// countBitsAndAppendBit counts the number of bits and appends the bit to the output.
func countBitsAndKeepBit(matrix [][]int, compareFn func(countZero int, countOne int) int) []int {
	// create the output
	output := make([]int, len(matrix))

	// find the most common bits
	for i := range matrix {
		countZero, countOne := countBits(matrix[i])

		// append the resulting bit
		output[i] = compareFn(countZero, countOne)
	}

	return output
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

// filterByMostCommonBits filters the matrix by the most common bits.
func filterByMostCommonBits(matrix [][]int) []int {
	compareFn := func(countZero int, countOne int) int {
		if countZero <= countOne {
			return 1
		} else {
			return 0
		}
	}

	return countBitsAndKeepBit(matrix, compareFn)
}

// filterByLeastCommonBits filters the matrix by the least common bits.
func filterByLeastCommonBits(matrix [][]int) []int {
	compareFn := func(countZero int, countOne int) int {
		if countZero > countOne {
			return 1
		} else {
			return 0
		}
	}

	return countBitsAndKeepBit(matrix, compareFn)
}

// convertArrayToString converts an array of ints to a string.
func convertArrayToString(array []int) string {
	// create the string
	str := make([]string, len(array))

	// fill the string
	for i := range array {
		str[i] = strconv.Itoa(array[i])
	}

	return strings.Join(str, "")
}

// filterMatrixByArray filters the matrix by the given array.
func filterMatrixByBin(matrix [][]int, filterFn func(matrix [][]int) []int) string {
	// create the output
	filteredMatrix := matrix

	// get the filter array
	filterArr := filterFn(transposeMatrix(filteredMatrix))

	// filter the matrix
	for i := range filterArr {
		var newFilteredMatrix [][]int

		if len(filteredMatrix) == 1 {
			// if there's only one value in the matrix, just return it
			break
		}

		for j := range filteredMatrix {
			if filterArr[i] == filteredMatrix[j][i] {
				newFilteredMatrix = append(newFilteredMatrix, filteredMatrix[j])
			}
		}

		// update the filtered matrix
		filteredMatrix = newFilteredMatrix

		// calculate the new filter array
		filterArr = filterFn(transposeMatrix(filteredMatrix))
	}

	return convertArrayToString(filteredMatrix[0])
}

// findParameters finds the gamma, epsilon, oxygen generator and the CO2
// scrubber ratings.
func findParameters(txtlines []string) (int, int, int, int) {
	var gammaBin, epsilonBin string
	var gamma, epsilon int64
	var oxygenGeneratorBin, CO2ScrubberBin string
	var oxygenGenerator, CO2Scrubber int64

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

	// find the oxygen generator rating
	oxygenGeneratorBin = filterMatrixByBin(matrix, filterByMostCommonBits)
	oxygenGenerator, _ = strconv.ParseInt(oxygenGeneratorBin, 2, 64)

	// find the CO2 scrubber rating
	CO2ScrubberBin = filterMatrixByBin(matrix, filterByLeastCommonBits)
	CO2Scrubber, _ = strconv.ParseInt(CO2ScrubberBin, 2, 64)

	return int(gamma), int(epsilon), int(oxygenGenerator), int(CO2Scrubber)
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// find the gamma rate, the epsilon rate, the oxygen generator rating, and
	// the CO2 scrubber rating
	gamma, epsilon, oxygenGenerator, CO2Scrubber := findParameters(txtlines)

	// calculate the power consumption
	power := gamma * epsilon

	// print the results
	fmt.Printf("Gamma: %d\n", gamma)
	fmt.Printf("Epsilon: %d\n", epsilon)
	fmt.Printf("Power: %d\n", power)

	// calculate the life support rating
	lifeSupport := oxygenGenerator * CO2Scrubber

	// print the results
	fmt.Printf("Oxygen Generator: %d\n", oxygenGenerator)
	fmt.Printf("CO2 Scrubber: %d\n", CO2Scrubber)
	fmt.Printf("Life Support: %d\n", lifeSupport)
}
