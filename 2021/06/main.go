package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// getInitialState returns the initial state of the fish
func getInitialState(txtlines []string) []int {
	var initialState []int

	for _, line := range txtlines {
		// convert the line to an int array
		if line != "" {
			states := strings.Split(line, ",")

			for _, state := range states {
				stateInt, _ := strconv.Atoi(state)
				initialState = append(initialState, stateInt)
			}
		}
	}

	return initialState
}

// getInitialFish returns the initial fish
func getInitialFish(initialState []int) []*LanternFish {
	initialFish := make([]*LanternFish, len(initialState))

	for i, state := range initialState {
		initialFish[i] = &LanternFish{}
		initialFish[i].new(state)
	}

	return initialFish
}

// getFishState returns the state of the fish
func getFishState(fish []*LanternFish) []int {
	fishState := make([]int, len(fish))

	for i, fish := range fish {
		fishState[i] = fish.getDaysLeft()
	}

	return fishState
}

// getFishAfterDays returns the fish after a set number of days
func getFishAfterDays(initialFish []*LanternFish, daysToCount int) []*LanternFish {
	fishAfterDays := initialFish

	for day := 1; day <= daysToCount; day++ {
		// loop through the fish
		for _, fish := range fishAfterDays {
			// get the next day fish
			nextDayFish := fish.getNextDayFish()

			// if the next day fish is not nil, append it to the list
			if nextDayFish != nil {
				fishAfterDays = append(fishAfterDays, nextDayFish)
			}
		}

		// print the number of fish after the current day
		if day < 20 && len(fishAfterDays) < 30 {
			dayWord := "days"

			if day == 1 {
				dayWord = "day"
			}

			currentState := getFishState(fishAfterDays)
			fmt.Printf("After %2d %s:\t%s\n", day, dayWord, helpers.IntArrayToString(currentState, ","))
		}
	}

	return fishAfterDays
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// get the initial state
	initialState := getInitialState(txtlines)

	// print the initial state
	fmt.Printf("Initial state:\t%s\n", helpers.IntArrayToString(initialState, ","))

	// create the initial fish
	initialFish := getInitialFish(initialState)

	// get the fish after a set number of days
	daysToCount := 80
	fishAfterDays := getFishAfterDays(initialFish, daysToCount)

	// print the text lines
	fmt.Printf("\nFinal number of fish: %d\n", len(fishAfterDays))
}
