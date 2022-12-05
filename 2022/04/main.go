package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const ELF_SEPARATOR = ","
const SECTION_SEPARATOR = "-"

// Section represents a section of the cleaning area [start, end]
type Section []int

// isPartiallyOverlapping checks if the sections are overlapping at all
func isPartiallyOverlapping(sections []Section) bool {
	sectionA := sections[0]
	sectionB := sections[1]

	sectionAStart := sectionA[0]
	sectionAEnd := sectionA[1]

	sectionBStart := sectionB[0]
	sectionBEnd := sectionB[1]

	if sectionAEnd >= sectionBStart && sectionBEnd >= sectionAStart {
		return true
	}

	return false
}

// isFullyOverlapping checks if the sections are fully overlapping
func isFullyOverlapping(sections []Section) bool {
	sectionA := sections[0]
	sectionB := sections[1]

	sectionAStart := sectionA[0]
	sectionAEnd := sectionA[1]

	sectionBStart := sectionB[0]
	sectionBEnd := sectionB[1]

	if sectionAStart >= sectionBStart && sectionAEnd <= sectionBEnd {
		return true
	}

	if sectionBStart >= sectionAStart && sectionBEnd <= sectionAEnd {
		return true
	}

	return false
}

// findPartiallyOverlappingSections finds the overlapping sections
func findPartiallyOverlappingSections(cleaningSections [][]Section) [][]Section {
	var overlappingSections [][]Section

	for _, sections := range cleaningSections {
		if isPartiallyOverlapping(sections) {
			overlappingSections = append(overlappingSections, sections)
		}
	}

	return overlappingSections
}

// findFullyOverlappingSections finds the fully overlapping sections
func findFullyOverlappingSections(cleaningSections [][]Section) [][]Section {
	var overlappingSections [][]Section

	for _, sections := range cleaningSections {
		if isFullyOverlapping(sections) {
			overlappingSections = append(overlappingSections, sections)
		}
	}

	return overlappingSections
}

// getCleaningSectionsFromInput returns the cleaning sections from the input
func getCleaningSectionsFromInput(input []string) [][]Section {
	var cleaningSections [][]Section

	for _, line := range input {
		sections := strings.Split(line, ELF_SEPARATOR)

		var cleaningSection []Section

		for _, section := range sections {
			sectionNumbers := strings.Split(section, SECTION_SEPARATOR)
			start, _ := strconv.Atoi(sectionNumbers[0])
			end, _ := strconv.Atoi(sectionNumbers[1])

			cleaningSection = append(cleaningSection, Section{start, end})
		}

		cleaningSections = append(cleaningSections, cleaningSection)
	}

	return cleaningSections
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	cleaningSections := getCleaningSectionsFromInput(txtlines)

	// part 1
	fullyOverlappingSections := findFullyOverlappingSections(cleaningSections)
	fullyOverlappingPairs := len(fullyOverlappingSections)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		fullyOverlappingPairs,
	)

	// part 2
	overlappingSections := findPartiallyOverlappingSections(cleaningSections)
	overlappingPairs := len(overlappingSections)
	fmt.Printf(
		"[Part Two] The answer is: %d\n",
		overlappingPairs,
	)
}
