package main

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/joaocarmo/advent-of-code/helpers"
)

// ALPHABET is the english alphabet.
const ALPHABET = "abcdefghijklmnopqrstuvwxyz"

// createRuneToIntoMap creates a map of runes to integers.
func createRuneToIntoMap() map[rune]int {
	counterStart := 1
	runeToInt := make(map[rune]int)

	// create the map for the lowercase letters
	for i, r := range ALPHABET {
		runeToInt[r] = counterStart + i
	}

	// create the map for the uppercase letters
	midStart := len(ALPHABET)
	for i, r := range ALPHABET {
		runeToInt[unicode.ToUpper(r)] = counterStart + midStart + i
	}

	return runeToInt
}

// calculateItemPriority calculates the priority of an item.
func calculateItemPriority(itemMap map[rune]int, item string) int {
	priority := 0

	for _, item := range item {
		priority += itemMap[item]
	}

	return priority
}

// calculateItemsPriority calculates the priority of a list of items.
func calculateItemsPriority(itemMap map[rune]int, items []string) []int {
	itemsPriority := []int{}

	for _, item := range items {
		itemsPriority = append(itemsPriority, calculateItemPriority(itemMap, item))
	}

	return itemsPriority
}

// findCommongItemsInArrays finds the common items in two arrays.
func findCommongItemsInArrays(arr1 []string, arr2 []string) []string {
	// create a map to store the items found
	commonItems := make(map[string]bool)

	// find the common items
	for _, item := range arr1 {
		if !commonItems[item] && helpers.StrArrayContains(arr2, item) {
			commonItems[item] = true
		}
	}

	// convert the map to a slice
	commonItemsSlice := []string{}
	for item := range commonItems {
		commonItemsSlice = append(commonItemsSlice, item)
	}

	return commonItemsSlice
}

// findCommonItemsInRucksack finds the common items in a rucksack.
func findCommonItemsInRucksack(items string) []string {
	// split the string in half
	half := len(items) / 2
	firstHalf := strings.Split(items[:half], "")
	secondHalf := strings.Split(items[half:], "")

	// convert the map to a slice
	commonItemsSlice := findCommongItemsInArrays(firstHalf, secondHalf)

	return commonItemsSlice
}

// findCommonItemInRucksacks finds the common item in a list of rucksacks.
func findCommonItemInRucksacks(items []string) []string {
	commonItemInRucksacks := []string{}

	for _, item := range items {
		commonItems := findCommonItemsInRucksack(item)
		commonItemInRucksacks = append(commonItemInRucksacks, commonItems[0])
	}

	return commonItemInRucksacks
}

// findCommonItemInRucksacksPerGroup finds the common item in a list of rucksacks.
func findCommonItemInRucksacksPerGroup(items []string) []string {
	groupSize := 3
	commonItemInRucksacks := []string{}

	for i := 0; i < len(items); i += groupSize {
		groupItems := items[i : i+groupSize]

		commonItems := [][]string{}
		for j := 0; j < groupSize - 1; j += 1 {
			firstGroupItems := strings.Split(groupItems[j], "")
			secondGroupItems := strings.Split(groupItems[j+1], "")
			commonItems = append(commonItems, findCommongItemsInArrays(firstGroupItems, secondGroupItems))
		}

		commonItemsInRucksacks := findCommongItemsInArrays(commonItems[0], commonItems[1])

		commonItemInRucksacks = append(commonItemInRucksacks, commonItemsInRucksacks[0])
	}

	return commonItemInRucksacks
}

// calculateTotalPriorities calculates the total priority of a list of priorities.
func calculateTotalPriorities(priorities []int) int {
	totalPriority := 0

	for _, priority := range priorities {
		totalPriority += priority
	}

	return totalPriority
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// create the map of runes to integers
	runeToInt := createRuneToIntoMap()

	// part 1
	commonItems := findCommonItemInRucksacks(txtlines)
	commonItemsPriority := calculateItemsPriority(runeToInt, commonItems)
	totalPriority := calculateTotalPriorities(commonItemsPriority)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		totalPriority,
	)

	// part 2
	commonItemsPerGroup := findCommonItemInRucksacksPerGroup(txtlines)
	commonItemsPerGroupPriority := calculateItemsPriority(runeToInt, commonItemsPerGroup)
	totalPerGroupPriority := calculateTotalPriorities(commonItemsPerGroupPriority)
	fmt.Printf(
		"[Part Two] The answer is: %d\n",
		totalPerGroupPriority,
	)
}
