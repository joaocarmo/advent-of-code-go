package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// stringToIntArray converts a string to an array of ints.
func stringToIntArray(str string, separator string) []int {
	var result []int

	for _, numStr := range strings.Split(str, separator) {
		cleanNumStr := strings.TrimSpace(numStr)

		if cleanNumStr != "" {
			num, _ := strconv.Atoi(cleanNumStr)
			result = append(result, num)
		}
	}

	return result
}

// cardIsEmpty checks if a card is empty.
func cardIsEmpty(card [][]int) bool {
	if len(card) > 0 && len(card[0]) > 0 {
		return false
	}

	return true
}

// bingoParse parses the text lines to get the random sequence and the bingo
// cards.
func bingoParse(txtlines []string) ([]int, [][][]int) {
	var sequence []int
	var cards [][][]int
	var card [][]int

	// making sure we have an extra stop condition for the loop
	stopCondition := ""
	txtlines = append(txtlines, stopCondition)

	for i, line := range txtlines {
		if i == 0 {
			sequence = stringToIntArray(line, ",")
			continue
		}

		if line == stopCondition {
			if !cardIsEmpty(card) {
				cards = append(cards, card)
				card = nil
			}
			continue
		}

		cardRow := stringToIntArray(line, " ")
		card = append(card, cardRow)
	}

	return sequence, cards
}

// findSmallestWinningSequence finds the winning card with the smallest winning
// sequence.
func findSmallestWinningSequence(winningCards []BingoCard) BingoCard {
	var winningCard BingoCard

	for _, card := range winningCards {
		if len(winningCard.winningSequence) == 0 || len(card.winningSequence) < len(winningCard.winningSequence) {
			winningCard = card
		}
	}

	return winningCard
}

// findWinningCard finds the winning card.
func findWinningCard(bingo []int, cards [][][]int) BingoCard {
	var winningCards []BingoCard

	for _, card := range cards {
		b := BingoCard{}

		b.new(bingo, card)

		if b.isWinner() {
			winningCards = append(winningCards, b)
		}
	}

	return findSmallestWinningSequence(winningCards)
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// parses the file for the random sequence and the bingo cards
	bingo, cards := bingoParse(txtlines)

	// find the winning card
	winningBingoCard := findWinningCard(bingo, cards)

	// print the winning card
	fmt.Println("winning card:\n")
	for _, row := range winningBingoCard.getCard() {
		for _, num := range row {
			fmt.Printf("%2d ", num)
		}
		fmt.Println()
	}
	fmt.Println()

	// print the winning sequence
	fmt.Printf("winning sequence: ")
	for _, num := range winningBingoCard.getWinningSequence() {
		fmt.Printf("%2d ", num)
	}
	fmt.Println("\n")

	// print the final score
	fmt.Printf("final score: %d\n", winningBingoCard.getScore())
}
