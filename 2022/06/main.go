package main

import (
	"fmt"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const startOfPacketMarker = 4

// pop removes the nth element from a slice
func pop(i int, xs []rune) (rune, []rune) {
  y := xs[i]
  ys := append(xs[:i], xs[i+1:]...)
  return y, ys
}

// getCharactersBeforePacketMarker returns the number of characters before the packet marker
func getCharactersBeforePacketMarker(message string) int {
	var charactersSeen []rune
	var numCharactersSeen int

	for count, character := range message {
		if helpers.RuneArrayContains(charactersSeen, character) {
			// remove the elements up to it from the list
			var removedCharacter rune
			var newList []rune
			for removedCharacter != character && len(newList) != startOfPacketMarker {
				removedCharacter, newList = pop(0, charactersSeen)
				charactersSeen = newList
			}
		}

		charactersSeen = append(charactersSeen, character)

		if len(charactersSeen) == startOfPacketMarker {
			numCharactersSeen = count
			break
		}
	}

	return numCharactersSeen + 1
}

// getCharactersBeforePacketMarkers returns the number of characters before the packet marker for each line
func getCharactersBeforePacketMarkers(messages []string) []int {
	var charactersBeforePacketMarker []int

	for _, message := range messages {
		charactersBeforePacketMarker = append(charactersBeforePacketMarker, getCharactersBeforePacketMarker(message))
	}

	return charactersBeforePacketMarker
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// part 1
	charactersBeforePacketMarker := getCharactersBeforePacketMarkers(txtlines)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		charactersBeforePacketMarker[0],
	)
}
