package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// comparePrevToCurrent compares the previous result to the current result.
func comparePrevToCurrent(prev, current int) int {
	if prev > current {
		return -1
	}

	if prev < current {
		return 1
	}

	return 0
}

// getStringFromResult returns the string representation of the result.
func getStringFromResult(result int) string {
	if result == -1 {
		return "decreased"
	}

	if result == 1 {
		return "increased"
	}

	return "no change"
}

// sumArrayElements sums the elements of an array.
func sumArrayElements(array []int) int {
	sum := 0

	for _, each := range array {
		sum += each
	}

	return sum
}

// getFinalAnswer returns the total number of times the result increased.
func getFinalAnswerPartOne(txtlines []string) int {
	var answer int
	var prev int
	var result int
	var resultStr string

	for _, eachline := range txtlines {
		// we'll convert each line from a string to an integer
		num, err := strconv.Atoi(eachline)

		if err != nil {
			log.Fatalf("failed converting string to integer: %s", err)
		}

		// compare the previous result to the current result
		if prev != 0 {
			result = comparePrevToCurrent(prev, num)
			resultStr = getStringFromResult(result)

			if result == 1 {
				answer++
			}
		} else {
			resultStr = "N/A - no previous measurement"
		}

		prev = num

		fmt.Printf("%d\t(%s)\n", num, resultStr)
	}

	return answer
}

// getFinalAnswerPartTwo returns the total number of times the sum of three
// consecutive measures increased.
func getFinalAnswerPartTwo(txtlines []string) int {
	var answer int
	var prev int
	var accum []int
	var sum int
	var result int
	var resultStr string

	for _, eachline := range txtlines {
		// we'll convert each line from a string to an integer
		num, err := strconv.Atoi(eachline)

		if err != nil {
			log.Fatalf("failed converting string to integer: %s", err)
		}

		// we need to accumulate the previous three measurements
		if len(accum) < 3 {
			accum = append(accum, num)
		} else {
			// we need to remove the first measurement and add the new one
			accum = append(accum[1:], num)
		}

		// if we have three measurements, we can calculate the sum
		if len(accum) == 3 {
			sum = sumArrayElements(accum)
		} else {
			continue
		}

		// compare the previous result to the current result
		if prev != 0 {
			result = comparePrevToCurrent(prev, sum)
			resultStr = getStringFromResult(result)

			if result == 1 {
				answer++
			}
		} else {
			resultStr = "N/A - no previous sum"
		}

		prev = sum

		fmt.Printf("%d\t(%s)\n", sum, resultStr)
	}

	return answer
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// get the final answers
	answerPartOne := getFinalAnswerPartOne(txtlines)

	fmt.Printf("[Part One] The answer is: %d\n", answerPartOne)

	answerPartTwo := getFinalAnswerPartTwo(txtlines)

	fmt.Printf("[Part Two] The answer is: %d\n", answerPartTwo)
}
