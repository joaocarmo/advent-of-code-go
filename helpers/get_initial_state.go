package helpers

import (
	"strconv"
	"strings"
)

// GetInitialState returns the initial state of the fish
func GetInitialState(txtlines []string) []int {
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
