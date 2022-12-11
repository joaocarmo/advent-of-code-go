package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = false
const NUM_OF_LINES_PER_MONKEY = 6
const STARTING_ITEMS = "Starting items: "
const STARTING_ITEMS_DELIMITER = ", "
const OLD_VALUE = "old"
const NUM_OF_ROUNDS_PART_1 = 20
const NUM_OF_ROUNDS_PART_2 = 1000
const RELIEF_DIVISOR_PART_1 = 3
const RELIEF_DIVISOR_PART_2 = 1
const NUM_MOST_ACTIVE_MONKEYS = 2

// Operation is an enum that represents the operation.
type Operation int

const (
	ADD Operation = iota
	MULTIPLY
)

// String returns the string representation of the operation.
func (o Operation) String() string {
	switch o {
	case ADD:
		return "+"
	case MULTIPLY:
		return "*"
	}
	return ""
}

// getOperationFromString returns the operation from a string.
func getOperationFromString(operation string) Operation {
	switch operation {
	case "+":
		return ADD
	case "*":
		return MULTIPLY
	}
	return -1
}

// OperationFn is a function that takes a worry level and returns a new worry level.
type OperationFn func (worryLevel int) int

// TestFn is a function that takes a worry level and returns a boolean.
type TestFn func (worryLevel int) bool

// IfConditionFn is a function that returns a monkey number.
type IfConditionFn func () int

// Monkey represents a monkey.
type Monkey struct {
	StartingItems  []int
	Operation      OperationFn
	Test           TestFn
	IfTrue         IfConditionFn
	IfFalse        IfConditionFn
	ItemsInspected int
}

// String returns the string representation of the monkey.
func (m Monkey) String() string {
	startingItemsString := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(m.StartingItems)), STARTING_ITEMS_DELIMITER), "[]")

	return startingItemsString
}

// inspect inspects the worry level item and returns a new worry level.
func (m *Monkey) inspect(item int) int {
	m.ItemsInspected++

	return m.Operation(item)
}

// play plays the game for a single monkey.
func (m *Monkey) play(monkeys []*Monkey, reliefDivisor int) {
	for _, item := range m.StartingItems {
		// Monkey inspects the worry level item
		newWorryLevel := m.inspect(item)
		// Monkey gets bored
		adjustedNewWorryLevel := newWorryLevel / reliefDivisor
		// Monkey tests the worry level item
		var throwToMonkey int
		if m.Test(adjustedNewWorryLevel) {
			throwToMonkey = m.IfTrue()
		} else {
			throwToMonkey = m.IfFalse()
		}

		if VERBOSE {
			fmt.Println("-- ", item, newWorryLevel, adjustedNewWorryLevel, throwToMonkey)
		}

		// Monkey throws the worry level item to another monkey
		monkeys[throwToMonkey].StartingItems = append(monkeys[throwToMonkey].StartingItems, adjustedNewWorryLevel)
		// Monkey gets rid of the worry level item
		m.StartingItems = m.StartingItems[1:]
	}
}

// getMonkeyNum returns the monkey number from a line.
func getMonkeyNum(line string) int {
	re := regexp.MustCompile(`Monkey (\d+):`)
	matches := re.FindStringSubmatch(line)
	monkeyNum, _ := strconv.Atoi(matches[1])

	return monkeyNum
}

// getStartingItems returns the starting items from a line.
func getStartingItems(line string) []int {
	startingItemsSlice := strings.Split(line, STARTING_ITEMS)
	startingItemsString := startingItemsSlice[1]
	startingItemsStringSlice := strings.Split(startingItemsString, STARTING_ITEMS_DELIMITER)
	startingItems := make([]int, len(startingItemsStringSlice))

	for i, startingItemString := range startingItemsStringSlice {
		startingItem, _ := strconv.Atoi(startingItemString)
		startingItems[i] = startingItem
	}

	return startingItems
}

// getOperation returns the operation from a line.
func getOperation(line string) OperationFn {
	re := regexp.MustCompile(`Operation: new = old ([\+|\*]) (\d+|\w+)`)
	matches := re.FindStringSubmatch(line)
	operation := getOperationFromString(matches[1])
	operandString := matches[2]
	operand, _ := strconv.Atoi(operandString)

	return func (worryLevel int) int {
		if operandString == OLD_VALUE {
			operand = worryLevel
		}

		switch operation {
		case ADD:
			return worryLevel + operand
		case MULTIPLY:
			return worryLevel * operand
		}

		return worryLevel
	}
}

// getTest returns the test from a line.
func getTest(line string) TestFn {
	re := regexp.MustCompile(`Test: divisible by (\d+)`)
	matches := re.FindStringSubmatch(line)
	divisible, _ := strconv.Atoi(matches[1])

	return func (worryLevel int) bool {
		return worryLevel % divisible == 0
	}
}

// getIfCondition returns the if condition from a line.
func getIfCondition(line string) IfConditionFn {
	re := regexp.MustCompile(`If (true|false): throw to monkey (\d+)`)
	matches := re.FindStringSubmatch(line)
	monkey, _ := strconv.Atoi(matches[2])

	return func () int {
		return monkey
	}
}

// calculateMonkeyBusiness calculates the monkey business.
func calculateMonkeyBusiness(monkeys []*Monkey) int {
	total := 1

	for _, monkey := range monkeys {
		total *= monkey.ItemsInspected
	}

	return total
}

// getMostActiveMonkeys returns the most active monkeys.
func getMostActiveMonkeys(monkeys []*Monkey) []*Monkey {
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].ItemsInspected > monkeys[j].ItemsInspected
	})

	return monkeys[:NUM_MOST_ACTIVE_MONKEYS]
}

// playKeepAway plays the game.
func playKeepAway(monkeys []*Monkey, numOfRounds, reliefDivisor int) {
	for i := 0; i < numOfRounds; i++ {
		if VERBOSE {
			fmt.Println("Round", i + 1)
		}

		for j, monkey := range monkeys {
			if VERBOSE {
				fmt.Println("- Monkey", j)
			}

			monkey.play(monkeys, reliefDivisor)
		}
	}
}

// getMonkeysFromFile returns the monkeys from a file.
func getMonkeysFromFile(txtlines []string) []*Monkey {
	numLines := len(txtlines)
	// get the number of monkeys, 1 is added to the number of lines to account for the empty line at the end of each monkey
	numLinesToProcess := NUM_OF_LINES_PER_MONKEY + 1
	numMonkeys := int(math.Round(float64(numLines) / float64(numLinesToProcess)))
	monkeys := make([]*Monkey, numMonkeys)

	// get the monkeys
	for i := 0; i < numLines; i += numLinesToProcess {
		// get the starting items
		monkeyNum := getMonkeyNum(txtlines[i])
		// get the starting items
		startingItems := getStartingItems(txtlines[i + 1])
		// get the operation
		operation := getOperation(txtlines[i + 2])
		// get the test
		test := getTest(txtlines[i + 3])
		// get the if true
		ifTrue := getIfCondition(txtlines[i + 4])
		// get the if false
		ifFalse := getIfCondition(txtlines[i + 5])

		// create the monkey
		monkey := Monkey{
			StartingItems: startingItems,
			Operation:     operation,
			Test:          test,
			IfTrue:        ifTrue,
			IfFalse:       ifFalse,
		}

		// add the monkey to the list
		monkeys[monkeyNum] = &monkey
	}

	return monkeys
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	monkeysA := getMonkeysFromFile(txtlines)

	// play the game
	playKeepAway(monkeysA, NUM_OF_ROUNDS_PART_1, RELIEF_DIVISOR_PART_1)

	if VERBOSE {
		for i, monkey := range monkeysA {
			fmt.Println(
				"Monkey",
				i,
				"has",
				len(monkey.StartingItems),
				"items: [",
				monkey,
				"], inspected",
				monkey.ItemsInspected,
				"times",
			)
		}
	}

	// part 1
	mostActiveMonkeys := getMostActiveMonkeys(monkeysA)
	monkeyBusiness := calculateMonkeyBusiness(mostActiveMonkeys)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		monkeyBusiness,
	)
}
