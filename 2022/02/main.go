package main

import (
	"fmt"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const WHITESPACE = " "

// Shape is a type of shape for the game (enum)
type Shape int

// Result is a type of result for the game (enum)
type Result int

const (
	Rock Shape = iota
	Paper
	Scissors
)

const (
	Draw Result = iota
	Win
	Lose
)

func (s Shape) String() string {
	switch s {
	case Rock:
		return "Rock"
	case Paper:
		return "Paper"
	case Scissors:
		return "Scissors"
	}
	return ""
}

func (s Shape) Int() int {
	switch s {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

func (r Result) String() string {
	switch r {
	case Draw:
		return "Draw"
	case Win:
		return "Win"
	case Lose:
		return "Lose"
	}
	return ""
}

func (r Result) Int() int {
	switch r {
	case Draw:
		return 3
	case Win:
		return 6
	case Lose:
		return 0
	}
	return 0
}

// calculateScore calculates the score for a chosen shape
func calculateScore(response Shape) int {
	switch response {
	case Rock:
		return 1
	case Paper:
		return 2
	case Scissors:
		return 3
	}
	return 0
}

// calculateOutcome calculates the outcome of a round
func calculateOutcome(opponent Shape, response Shape) int {
	if opponent == response {
		return Draw.Int()
	}

	switch opponent {
	case Rock:
		if response == Paper {
			return Win.Int()
		}
		return Lose.Int()
	case Paper:
		if response == Scissors {
			return Win.Int()
		}
		return Lose.Int()
	case Scissors:
		if response == Rock {
			return Win.Int()
		}
		return Lose.Int()
	}

	return 0
}

// calculateRoundScore calculates the score for a round
func calculateRoundScore(opponent Shape, response Shape) int {
	scoreForChosenShape := calculateScore(response)
	scoreForRoundOutcome := calculateOutcome(opponent, response)

	return scoreForChosenShape + scoreForRoundOutcome
}

// calculateTotalScore calculates the total score for the game
func calculateTotalScore(shapes [][]Shape) int {
	totalScore := 0

	for _, shape := range shapes {
		totalScore += calculateRoundScore(shape[0], shape[1])
	}

	return totalScore
}

// convertOpponentToShape converts the opponent's response to a shape
func convertOpponentToShape(input string) Shape {
	switch input {
	case "A":
		return Rock
	case "B":
		return Paper
	case "C":
		return Scissors
	}
	return 0
}

// convertResponseToShape converts the response to a shape
func convertResponseToShape(input string) Shape {
	switch input {
	case "X":
		return Rock
	case "Y":
		return Paper
	case "Z":
		return Scissors
	}
	return 0
}

// convertInputToShapes converts the input to shapes
func convertInputToShapes(input []string) [][]Shape {
	shapes := make([][]Shape, len(input))

	for i, line := range input {
		// split the line by a white space into two parts
		strategy := strings.Split(line, WHITESPACE)

		shapes[i] = make([]Shape, 2)
		shapes[i][0] = convertOpponentToShape(string(strategy[0]))
		shapes[i][1] = convertResponseToShape(string(strategy[1]))
	}

	return shapes
}

// main is the entry point for the application
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	shapes := convertInputToShapes(txtlines)

	// part 1
	totalScore := calculateTotalScore(shapes)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		totalScore,
	)
}
