package main

import (
	"fmt"
	"strconv"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// distributeSnacksEachElfIsCarrying distributes the snacks each elf is carrying
func distributeSnacksEachElfIsCarrying(input []string) [][]int {
	// Each elf has a list of snacks per line
	// Blank lines separate each elf
	var elfs [][]int
	var elf []int

	// Split the input into elfs
	for _, line := range input {
		snack, err := strconv.Atoi(line)

		if err != nil {
			// Blank line, new elf
			elfs = append(elfs, elf)
			elf = []int{}
		} else {
			elf = append(elf, snack)
		}
	}

	// Add the last elf, if not empty
	if len(elf) > 0 {
		elfs = append(elfs, elf)
	}

	return elfs
}

// calculateTotalNumberOfCaloriesEachElfIsCarrying calculates the total number of calories each elf is carrying
func calculateTotalNumberOfCaloriesEachElfIsCarrying(input []string) []int {
	snacksEachElfIsCarrying := distributeSnacksEachElfIsCarrying(input)

	var totalNumberOfCaloriesEachElfIsCarrying []int

	for _, elf := range snacksEachElfIsCarrying {
		var totalNumberOfCalories int

		for _, snack := range elf {
			totalNumberOfCalories += snack
		}

		totalNumberOfCaloriesEachElfIsCarrying = append(totalNumberOfCaloriesEachElfIsCarrying, totalNumberOfCalories)
	}

	return totalNumberOfCaloriesEachElfIsCarrying
}

// findTheElfWithTheMostCalories finds the elf with the most calories
func findTheElfWithTheMostCalories(input []int) (int, int) {
	var elfWithTheMostCalories int
	var mostCalories int

	for elf, calories := range input {
		if calories > mostCalories {
			elfWithTheMostCalories = elf
			mostCalories = calories
		}
	}

	return elfWithTheMostCalories, mostCalories
}

// findTheTopThreeElfsWithTheMostCalories finds the top three elfs with the most calories
func findTheTopThreeElfsWithTheMostCalories(input []int) ([]int, int) {
	var topThreeElfsWithTheMostCalories []int
	topThreeElfsMostCalories := 0

	// Copy the input
	elfsWithCalories := make([]int, len(input))
	copy(elfsWithCalories, input)

	// Find the first elf with the most calories
	elfWithTheMostCalories, mostCalories := findTheElfWithTheMostCalories(elfsWithCalories)
	topThreeElfsWithTheMostCalories = append(topThreeElfsWithTheMostCalories, elfWithTheMostCalories)
	topThreeElfsMostCalories += mostCalories

	// Remove the first elf from the list
	elfsWithCalories[elfWithTheMostCalories] = 0

	// Find the second elf with the most calories
	elfWithTheMostCalories, mostCalories = findTheElfWithTheMostCalories(elfsWithCalories)
	topThreeElfsWithTheMostCalories = append(topThreeElfsWithTheMostCalories, elfWithTheMostCalories)
	topThreeElfsMostCalories += mostCalories

	// Remove the second elf from the list
	elfsWithCalories[elfWithTheMostCalories] = 0

	// Find the third elf with the most calories
	elfWithTheMostCalories, mostCalories = findTheElfWithTheMostCalories(elfsWithCalories)
	topThreeElfsWithTheMostCalories = append(topThreeElfsWithTheMostCalories, elfWithTheMostCalories)
	topThreeElfsMostCalories += mostCalories

	return topThreeElfsWithTheMostCalories, topThreeElfsMostCalories
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	totalNumberOfCaloriesEachElfIsCarrying := calculateTotalNumberOfCaloriesEachElfIsCarrying(txtlines)

	// part 1
	elfWithTheMostCalories, mostCalories := findTheElfWithTheMostCalories(totalNumberOfCaloriesEachElfIsCarrying)
	fmt.Printf(
		"[Part One] The answer is: %7d\t(elf  #%d)\n",
		mostCalories,
		elfWithTheMostCalories + 1,
	)

	// part 2
	topThreeElfsWithTheMostCalories, topThreeElfsMostCalories := findTheTopThreeElfsWithTheMostCalories(totalNumberOfCaloriesEachElfIsCarrying)

	// print the text lines
	fmt.Printf(
		"[Part Two] The answer is: %7d\t(elfs #%d #%d #%d)\n",
		topThreeElfsMostCalories, topThreeElfsWithTheMostCalories[0] + 1,
		topThreeElfsWithTheMostCalories[1] + 1,
		topThreeElfsWithTheMostCalories[2] + 1,
	)
}
