package main

import (
	"fmt"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const verbose = false

type Procedure struct {
	Move int
	From int
	To   int
}

type Stacks map[int][]string

// getTopCratesFromStacks returns the top crates from the stacks.
func getTopCratesFromStacks(stacks Stacks) string {
	var topCrates string

	keysMap := reflect.ValueOf(stacks).MapKeys()
	keys := make([]int, len(keysMap))
	for i, key := range keysMap {
		keys[i] = key.Interface().(int)
	}
	sort.Ints(keys)

	for _, key := range keys {
		lastIndex := len(stacks[key]) - 1
		topCrates += stacks[key][lastIndex]
	}

	return topCrates
}

// arrangeMultipleStacksByProcedure arranges the stacks by a procedure (part 2).
func arrangeMultipleStacksByProcedure(stacks Stacks, procedure Procedure) Stacks {
	if verbose {
		fmt.Println("==============================================")
		fmt.Printf(" -> move %d from %d to %d\n", procedure.Move, procedure.From, procedure.To)
		fmt.Println("==============================================")
		for key, value := range stacks {
			fmt.Printf("stack %d: %s\n", key, strings.Join(value, ", "))
		}
		fmt.Println("----------------------------------------------")
	}

	fromStack := stacks[procedure.From]
	toStack := stacks[procedure.To]

	indexOfFirstCrateToMove := len(fromStack) - procedure.Move

	cratesToMove := fromStack[indexOfFirstCrateToMove:]

	// remove the crates from the `from` stack
	stacks[procedure.From] = fromStack[:indexOfFirstCrateToMove]

	// add the crates to the `to` stack
	stacks[procedure.To] = append(toStack, cratesToMove...)

	if verbose {
		for key, value := range stacks {
			fmt.Printf("stack %d: %s\n", key, strings.Join(value, ", "))
		}
	}

	return stacks
}

// arrangeMultipleStacks arranges the stacks by a slice of procedures (part 2).
func arrangeMultipleStacks(stacks Stacks, procedures []Procedure) Stacks {
	// copy the stacks
	arrangedStacks := make(Stacks)

	for key, value := range stacks {
		arrangedStacks[key] = value
	}

	// arrange the stacks
	for _, procedure := range procedures {
		arrangedStacks = arrangeMultipleStacksByProcedure(arrangedStacks, procedure)
	}

	return arrangedStacks
}

// arrangeStacksByProcedure arranges the stacks by a procedure.
func arrangeStacksByProcedure(stacks Stacks, procedure Procedure) Stacks {
	if procedure.Move < 1 {
		return stacks
	}

	// remove the crate from the `from` stack
	crateIndex := len(stacks[procedure.From]) - 1
	crate := stacks[procedure.From][crateIndex]
	stacks[procedure.From] = stacks[procedure.From][:crateIndex]

	// add the crate to the `to` stack
	stacks[procedure.To] = append(stacks[procedure.To], crate)

	// mark the crate as moved
	procedure.Move = procedure.Move - 1

	return arrangeStacksByProcedure(stacks, procedure)
}

// arrangeStacks arranges the stacks by a slice of procedures.
func arrangeStacks(stacks Stacks, procedures []Procedure) Stacks {
	// copy the stacks
	arrangedStacks := make(Stacks)

	for key, value := range stacks {
		arrangedStacks[key] = value
	}

	// arrange the stacks
	for _, procedure := range procedures {
		arrangedStacks = arrangeStacksByProcedure(arrangedStacks, procedure)
	}

	return arrangedStacks
}

// parseProcedure parses a procedure into a Procedure struct.
func parseProcedure(procedure string) Procedure {
	var parsedProcedure Procedure

	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	matches := re.FindStringSubmatch(procedure)

	parsedProcedure.Move, _ = strconv.Atoi(matches[1])
	parsedProcedure.From, _ = strconv.Atoi(matches[2])
	parsedProcedure.To, _ = strconv.Atoi(matches[3])

	return parsedProcedure
}

// parseProcedures parses the procedures into a slice of Procedure structs.
func parseProcedures(procedures []string) []Procedure {
	var parsedProcedures []Procedure

	for _, procedure := range procedures {
		parsedProcedure := parseProcedure(procedure)
		parsedProcedures = append(parsedProcedures, parsedProcedure)
	}

	return parsedProcedures
}

// parseStacks parses the stacks into a map of stacks.
func parseStacks(stacks []string) Stacks {
	// create map of stacks using numbers as keys
	stacksMap := make(Stacks)

	// copy the stacks
	var parsedStacks []string
	parsedStacks = append(parsedStacks, stacks...)

	// reverse the stacks
	for i, j := 0, len(parsedStacks)-1; i < j; i, j = i+1, j-1 {
		parsedStacks[i], parsedStacks[j] = parsedStacks[j], parsedStacks[i]
	}

	// parse the stacks matrix
	for i, value := range parsedStacks[0] {
		stackNumber, err := strconv.Atoi(string(value))

		if err == nil {
			for _, stackString := range parsedStacks[1:] {
				if len(stackString) <= i {
					continue
				}

				crate := strings.TrimSpace(string(stackString[i]))

				if crate != "" {
					stacksMap[stackNumber] = append(stacksMap[stackNumber], string(stackString[i]))
				}
			}
		}
	}

	return stacksMap
}

// getStacksAndProcedures returns the stacks and procedures from the input.
func getStacksAndProcedures(input []string) ([]string, []string) {
	var stacks []string
	var procedures []string

	stackIsDone := false
	for _, line := range input {
		if line == "" {
			stackIsDone = true
			continue
		}

		if stackIsDone {
			procedures = append(procedures, line)
		} else {
			stacks = append(stacks, line)
		}
	}

	return stacks, procedures
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	stacks, procedures := getStacksAndProcedures(txtlines)
	parsedStacks := parseStacks(stacks)
	parsedProcedures := parseProcedures(procedures)

	// part 1
	arrangedStacks := arrangeStacks(parsedStacks, parsedProcedures)
	topCrates := getTopCratesFromStacks(arrangedStacks)
	fmt.Printf(
		"[Part One] The answer is: %s\n",
		topCrates,
	)

	// part 2
	arrangedMultiStacks := arrangeMultipleStacks(parsedStacks, parsedProcedures)
	topMultiCrates := getTopCratesFromStacks(arrangedMultiStacks)
	fmt.Printf(
		"[Part Two] The answer is: %s\n",
		topMultiCrates,
	)
}
