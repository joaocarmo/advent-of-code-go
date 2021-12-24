package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const verbose = false

var numOfSignalsToDigit = map[int]int{
	0: 6,
	1: 2, // unique
	2: 5,
	3: 5,
	4: 4, // unique
	5: 5,
	6: 6,
	7: 3, // unique
	8: 7, // unique
	9: 6,
}

// getNumOfOccurrences returns the number of occurrences of a signal in an array
// of signals.
func getNumOfOccurrences(signal string, signals []string) int {
	count := 0

	for _, s := range signals {
		count += strings.Count(s, signal)
	}

	return count
}

// printOutputMap prints the output map.
func printOutputMap(output map[int]map[string][]int) {
	if verbose {
		line := 0

		// loop through the lines
		for _, lineDigits := range output {
			line += 1

			// loop through the output
			for signal, digits := range lineDigits {
				fmt.Printf("Line %-2d | %-8s: %v\n", line, signal, digits)
			}
		}
	}
}

// parseInput parses the input into a map of signal patterns and output values.
func parseInput(lines []string) ([][]string, [][]string) {
	allSignals := make([][]string, len(lines))
	allOutput := make([][]string, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}

		// split the line into signal and output
		result := strings.Split(line, "|")
		signalPatterns := strings.TrimSpace(result[0])
		outputValues := strings.TrimSpace(result[1])

		// split the signal patterns into an array
		signals := strings.Split(signalPatterns, " ")

		// split the output values into an array
		output := strings.Split(outputValues, " ")

		// add the signal and output to the map
		allSignals[i] = signals
		allOutput[i] = output
	}

	return allSignals, allOutput
}

// inferDigitsFromNumOfSignals takes a number of signals and returns the digit.
func inferDigitsFromNumOfSignals(signal string) []int {
	// create a map of the signals
	var possibleDigits []int

	// count the number of signals
	numSignals := len(signal)

	// loop through numOfSignalsToDigit
	for digit, numOfSignals := range numOfSignalsToDigit {
		// if the number of signals is the required number for the digit
		if numSignals == numOfSignals {
			// add the digit to the possible digits
			possibleDigits = append(possibleDigits, digit)
		}
	}

	// return the digit
	return possibleDigits
}

// getPossibleOutputDigits returns the possible output digits.
func getPossibleOutputDigits(output [][]string, collapse bool) map[int]map[string][]int {
	// create a map of the possible digits
	var possibleDigits = make(map[int]map[string][]int, len(output))

	// loop through the output
	for i, lines := range output {
		// loop through the digits
		for _, digit := range lines {
			if _, ok := possibleDigits[i]; !ok {
				// create the map
				possibleDigits[i] = make(map[string][]int)
			}

			var newKey string
			counter := 0

			for {

				if collapse {
					newKey = digit
				} else {
					newKey = fmt.Sprintf("%s_%d", digit, counter)
				}

				// if the digit is not in the map
				if _, ok := possibleDigits[i][newKey]; !ok {
					// add the digit to the map
					possibleDigits[i][newKey] = inferDigitsFromNumOfSignals(digit)
					break
				}

				counter += 1
			}
		}
	}

	// return the possible digits
	return possibleDigits
}

// getSingleDigits returns the single digits from the possible output digits.
func getSingleDigits(possibleOutputDigits map[int]map[string][]int) map[int]map[string][]int {
	// create a map of the single digits
	var singleDigits = make(map[int]map[string][]int)

	// loop through the lines
	for i, lineDigits := range possibleOutputDigits {
		// loop through the possible digits
		for signal, digits := range lineDigits {
			// if there is only one digit
			if len(digits) == 1 {
				if _, ok := singleDigits[i]; !ok {
					// create the map
					singleDigits[i] = make(map[string][]int)
				}

				// if the digit is not in the map
				if _, ok := singleDigits[i][signal]; !ok {
					// add the digit to the map
					singleDigits[i][signal] = digits
				}
			}
		}
	}

	// return the single digits
	return singleDigits
}

// countLenPerLine counts the number of unique signals per line.
func countLenPerLine(output map[int]map[string][]int) int {
	// create a map of the number of unique signals per line
	numOfUniqueSignals := 0

	// loop through the lines
	for _, lineDigits := range output {
		// loop through the digits
		for _, digits := range lineDigits {
			numOfUniqueSignals += len(digits)
		}
	}

	// return the number of unique signals
	return numOfUniqueSignals
}

// getSignalsContaining returns the signals containing a given possible number.
func getSignalsContaining(lineDigits map[string][]int, number int) []string {
	var signalsContaining []string

	// loop through the digits
	for signal, digits := range lineDigits {
		if helpers.IntArrayContains(digits, number) {
			signalsContaining = append(signalsContaining, signal)
		}
	}

	return signalsContaining
}

// getSevenSegmentDisplaysFromSignals gets the 7-segment display decoded from the
// signals.
func getSevenSegmentDisplaysFromSignals(signals [][]string) []*SevenSegmentDisplay {
	var ssdArr []*SevenSegmentDisplay
	singalsContaining := make(map[int][]string, 10)

	possibleOutputDigits := getPossibleOutputDigits(signals, true)
	printOutputMap(possibleOutputDigits)

	// loop through the lines
	for _, lineDigits := range possibleOutputDigits {
		for i := 0; i < 10; i++ {
			singalsContaining[i] = getSignalsContaining(lineDigits, i)

			if verbose {
				fmt.Println(i, singalsContaining[i])
			}
		}

		ssd := &SevenSegmentDisplay{}
		ssd.inferFromSingals(singalsContaining)
		ssdArr = append(ssdArr, ssd)
	}

	return ssdArr
}

// getDecodedOutput gets the decoded output from the signals.
func getDecodedOutput(output [][]string, ssdArr []*SevenSegmentDisplay) [][]int {
	var decodedOutput [][]int

	for line, signals := range output {
		decodedLine := []int{}

		for _, signal := range signals {
			ssd := ssdArr[line]
			decodedLine = append(decodedLine, ssd.getNumForSignal(signal))
		}

		decodedOutput = append(decodedOutput, decodedLine)
	}

	return decodedOutput
}

// decodedOutputToNumbers converts the decoded output to numbers.
func decodedOutputToNumbers(decodedOutput [][]int) []int {
	var numbers []int

	for _, digits := range decodedOutput {
		num, _ := strconv.Atoi(helpers.IntArrayToString(digits, ""))
		numbers = append(numbers, num)
	}

	return numbers
}

// sumDecodedNumbers sums the decoded numbers.
func sumDecodedNumbers(numbers []int) int {
	sum := 0

	for _, num := range numbers {
		sum += num
	}

	return sum
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// parse the file into signal patterns and output values
	signals, output := parseInput(txtlines)

	// infer the digits from the number of signals for the output
	possibleOutputDigits := getPossibleOutputDigits(output, false)

	// print information
	printOutputMap(possibleOutputDigits)
	fmt.Printf("Number of unique signals: %d\n\n", countLenPerLine(possibleOutputDigits))

	// get the number of single digits for a given signal
	numOfSingleDigits := getSingleDigits(possibleOutputDigits)

	// print information
	printOutputMap(numOfSingleDigits)
	fmt.Printf("Number of unique signals with single possible digits: %d\n\n", countLenPerLine(numOfSingleDigits))

	// get the 7-segment display from the singnals
	ssdArr := getSevenSegmentDisplaysFromSignals(signals)

	// get the decoded output from the 7-segment displays
	decodedOutput := getDecodedOutput(output, ssdArr)

	decodedNumbers := decodedOutputToNumbers(decodedOutput)

	sumOfNumbers := sumDecodedNumbers(decodedNumbers)

	// print the decoded output
	fmt.Printf("Decoded output: %v\n\n", decodedNumbers)
	fmt.Printf("Sum of output: %d\n", sumOfNumbers)

	// print the 7-segment displays
	if verbose {
		fmt.Println()
		for line, ssd := range ssdArr {
			fmt.Printf("Line %d, 7-segment display:\n%s\n", line, ssd.toString())
			for i := 0; i < 10; i++ {
				fmt.Printf("%d: %s\n", i, ssd.getSingalForNum(i))
			}
			fmt.Println()
		}
	}
}
