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

// isOverlapping checks if the sections are overlapping
func isOverlapping(sections []Section) bool {
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

// findOverlappingSections finds the overlapping sections
func findOverlappingSections(cleaningSections [][]Section) [][]Section {
	var overlappingSections [][]Section

	for _, sections := range cleaningSections {
		if isOverlapping(sections) {
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

	// part 1
	cleaningSections := getCleaningSectionsFromInput(txtlines)
	overlappingSections := findOverlappingSections(cleaningSections)
	overlappingPairs := len(overlappingSections)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		overlappingPairs,
	)
}
