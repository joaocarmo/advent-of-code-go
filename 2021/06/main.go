package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const useVersion = 2
const verbose = false
const dayThreshold = 20
const daysToCount = 256

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

// printStatus prints the current status of the fish
func printStatus(day int, fish []*LanternFish) {
	if !verbose {
		return
	}

	dayWord := "days"
	var currentState string

	if day == 1 {
		dayWord = "day"
	}
	if useVersion == 1 && day < dayThreshold && len(fish) < 30 {
		currentState = helpers.IntArrayToString(getFishState(fish), ",")
	} else {
		currentState = "..."
	}

	fmt.Printf("After %2d %s:\t%s\n", day, dayWord, currentState)
}

// getFishAfterDays returns the fish after a set number of days (version 1)
func getFishAfterDaysV1(initialFish []*LanternFish, daysToCount int) []*LanternFish {
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
		printStatus(day, fishAfterDays)
	}

	return fishAfterDays
}

// memoizedGetFishAfterDays returns the fish after a set number of days
func memoizedGetFishAfterNDays(cache map[string]int) func(lf *LanternFish, daysToCount int) int {
	return func(lf *LanternFish, daysToCount int) int {
		getFishAfterNDays := memoizedGetFishAfterNDays(cache)
		daysLeftKey := strconv.Itoa(lf.getDaysLeft())
		daysToCountKey := strconv.Itoa(daysToCount)
		cacheKey := daysLeftKey + "-" + daysToCountKey

		if cached, ok := cache[cacheKey]; ok {
			if verbose {
				fmt.Printf("Cache hit: %s\n", cacheKey)
			}

			return cached
		}

		fishAfterDays := 1

		// loop through the day
		for day := 1; day <= daysToCount; day++ {
			// get the next day fish
			nextDayFish := lf.getNextDayFish()

			// if the next day fish is not nil, append it to the list
			if nextDayFish != nil {
				fishAfterDays += getFishAfterNDays(nextDayFish, daysToCount-day)
			}
		}

		cache[cacheKey] = fishAfterDays

		if verbose {
			fmt.Printf("Cache miss: %s\n", cacheKey)
		}

		return cache[cacheKey]
	}
}

// getFishAfterDays returns the fish after a set number of days (version 2)
func getFishAfterDaysV2(initialFish []*LanternFish, daysToCount int) int {
	fishAfterDays := 0

	var cache = make(map[string]int)
	getFishAfterNDays := memoizedGetFishAfterNDays(cache)

	for _, fishDaysLeft := range initialFish {
		fishAfterDays += getFishAfterNDays(fishDaysLeft, daysToCount)
	}

	return fishAfterDays
}

// getFishAfterDays returns the fish after a set number of days
func getFishAfterDays(initialFish []*LanternFish, daysToCount int) int {
	if useVersion == 1 {
		return len(getFishAfterDaysV1(initialFish, daysToCount))
	}

	return getFishAfterDaysV2(initialFish, daysToCount)
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// get the initial state
	initialState := helpers.GetInitialState(txtlines)

	// print the initial state
	fmt.Printf("Initial state:\t%s\n", helpers.IntArrayToString(initialState, ","))

	// create the initial fish
	initialFish := getInitialFish(initialState)

	// get the fish after a set number of days
	start := time.Now()
	fishAfterDays := getFishAfterDays(initialFish, daysToCount)
	elapsed := time.Since(start)

	// print the text lines
	fmt.Printf("\nFinal number of fish: %d\n", fishAfterDays)

	// print the elapsed time
	fmt.Printf("\nElapsed time (v%d): %s\n", useVersion, elapsed)
}
