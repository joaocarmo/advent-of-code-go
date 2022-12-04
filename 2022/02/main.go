package main

import (
	"fmt"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const WHITESPACE = " "

// Shape is a type of shape for the game (enum)
type Shape int

// Outcome is a type of outcome for the game (enum)
type Outcome int

const (
	Rock Shape = iota
	Paper
	Scissors
)

const (
	Draw Outcome = iota
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

func (o Outcome) String() string {
	switch o {
	case Draw:
		return "Draw"
	case Win:
		return "Win"
	case Lose:
		return "Lose"
	}
	return ""
}

func (o Outcome) Int() int {
	switch o {
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

// calculateResponse calculates the response of a round
func calculateResponse(opponent Shape, outcome Outcome) Shape {
	if outcome == Draw {
		return opponent
	}

	switch outcome {
	case Win:
		switch opponent {
		case Rock:
			return Paper
		case Paper:
			return Scissors
		case Scissors:
			return Rock
		}
	case Lose:
		switch opponent {
		case Rock:
			return Scissors
		case Paper:
			return Rock
		case Scissors:
			return Paper
		}
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
func calculateTotalScore(s1 []Shape, s2 []Shape) int {
	totalScore := 0

	for i, shape := range s1 {
		totalScore += calculateRoundScore(shape, s2[i])
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

// convertResponseToShape converts the response to a shape (part 1)
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

// convertResponseToShape converts the response to a shape (part 2)
func convertResponseToOutcome(input string) Outcome {
	switch input {
	case "X":
		return Lose
	case "Y":
		return Draw
	case "Z":
		return Win
	}
	return 0
}

// convertShapesAndOutcomesToResponses converts the shapes and outcomes to responses
func convertShapesAndOutcomesToResponses(opponent []Shape, outcome []Outcome) []Shape {
	responses := make([]Shape, len(opponent))

	for i, shape := range opponent {
		responses[i] = calculateResponse(shape, outcome[i])
	}

	return responses
}

// convertInputToShapes converts the input to shapes (part 1)
func convertInputToShapes(input []string) ([]Shape, []Shape) {
	shapes := make([]Shape, len(input))
	responses := make([]Shape, len(input))

	for i, line := range input {
		// split the line by a white space into two parts
		strategy := strings.Split(line, WHITESPACE)

		shapes[i] = convertOpponentToShape(string(strategy[0]))
		responses[i] = convertResponseToShape(string(strategy[1]))
	}

	return shapes, responses
}

// convertInputToShapes converts the input to shapes (part 2)
func convertInputToShapesAndOutcomes(input []string) ([]Shape, []Outcome) {
	shapes := make([]Shape, len(input))
	outcomes := make([]Outcome, len(input))

	for i, line := range input {
		// split the line by a white space into two parts
		strategy := strings.Split(line, WHITESPACE)

		shapes[i] = convertOpponentToShape(string(strategy[0]))
		outcomes[i] = convertResponseToOutcome(string(strategy[1]))
	}

	return shapes, outcomes
}

// main is the entry point for the application
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	shapesPartOne, responsesPartOne := convertInputToShapes(txtlines)

	// part 1
	totalScorePartOne := calculateTotalScore(shapesPartOne, responsesPartOne)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		totalScorePartOne,
	)

	// part 2
	shapes, outcomes := convertInputToShapesAndOutcomes(txtlines)
	responses := convertShapesAndOutcomesToResponses(shapes, outcomes)
	totalScorePartTwo := calculateTotalScore(shapes, responses)
	fmt.Printf(
		"[Part Two] The answer is: %d\n",
		totalScorePartTwo,
	)
}
