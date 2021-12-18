package main

import (
	"fmt"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const verbose = true

// getCrabPositions parses the text lines to get the crab positions.
func getCrabPositions(txtlines []string) []*Crab {
	initialState := helpers.GetInitialState(txtlines)
	crabs := make([]*Crab, len(initialState))

	for i, initialPosition := range initialState {
		crabs[i] = &Crab{}
		crabs[i].new(initialPosition, 0, 0)
	}

	return crabs
}

// getLowerHigherLimit returns the lower and higher limit for the given crabs.
func getLowerHigherLimit(crabs []*Crab) (int, int) {
	if len(crabs) == 0 {
		return 0, 0
	}

	lowerLimit := crabs[0].getX()
	higherLimit := crabs[0].getX()

	for _, crab := range crabs {
		if crab.getX() < lowerLimit {
			lowerLimit = crab.getX()
		}

		if crab.getX() > higherLimit {
			higherLimit = crab.getX()
		}
	}

	return lowerLimit, higherLimit
}

// getDistance returns the distance between the given crabs and the given
// position.
func getDistance(crabs []*Crab, positionX int) int {
	var distance int

	for _, crab := range crabs {
		distance += crab.getDistanceTo(positionX, 0, 0)
	}

	return distance
}

// getOptimalPositionAndFuel returns the optimal position and fuel consumption
// for the crabs. Considering only the horizontal axis.
func getOptimalPositionAndFuel(crabs []*Crab) (int, int) {
	var optimalPosition int
	var fuelConsumption int

	lowLimit, highLimit := getLowerHigherLimit(crabs)

	for {
		thirds := (highLimit - lowLimit) / 3
		mid1 := lowLimit + thirds
		mid2 := highLimit - thirds

		dist1 := getDistance(crabs, mid1)
		dist2 := getDistance(crabs, mid2)

		if dist1 < dist2 {
			highLimit = mid2
		} else {
			lowLimit = mid1
		}

		if verbose {
			// print the results
			fmt.Printf("Low limit: %d\n", lowLimit)
			fmt.Printf("High limit: %d\n", highLimit)
			fmt.Printf("Distance1: %d\n", dist1)
			fmt.Printf("Distance2: %d\n", dist2)
		}

		if highLimit-lowLimit <= 2 {
			break
		}
	}

	optimalPosition = lowLimit + (highLimit-lowLimit)/2
	fuelConsumption = getDistance(crabs, optimalPosition)

	return optimalPosition, fuelConsumption
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// parse the file to get the crab positions
	crabs := getCrabPositions(txtlines)

	// get the optimal position
	optimalPosition, fuelConsumption := getOptimalPositionAndFuel(crabs)

	// print the results
	fmt.Printf("Optimal position: %d\n", optimalPosition)
	fmt.Printf("Fuel consumption: %d\n", fuelConsumption)
}
