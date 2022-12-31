package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = false
const ROOT = "root"

// Operator represents an operator.
type Operator int

const (
	Plus Operator = iota
	Minus
	Multiply
	Divide
)

// String returns the string representation of an operator.
func (o Operator) String() string {
	return [...]string{"+", "-", "*", "/"}[o]
}

// Job represents a job.
type Job int

const (
	YellNumber Job = iota
	YellOperation
)

// newOperator creates a new operator.
func newOperator(operator string) Operator {
	switch operator {
	case "+":
		return Plus
	case "-":
		return Minus
	case "*":
		return Multiply
	case "/":
		return Divide
	}

	return -1
}

// String returns the string representation of a job.
func (j Job) String() string {
	return [...]string{"YellNumber", "YellOperation"}[j]
}

// Operation represents an operation.
type Operation struct {
	operator  Operator
	leftSide  *Monkey
	rightSide *Monkey
}

// getResult returns the result of an operation.
func (o Operation) getResult() int {
	var result int

	switch o.operator {
	case Plus:
		result = o.leftSide.result + o.rightSide.result
	case Minus:
		result = o.leftSide.result - o.rightSide.result
	case Multiply:
		result = o.leftSide.result * o.rightSide.result
	case Divide:
		result = o.leftSide.result / o.rightSide.result
	}

	return result
}

// String returns the string representation of an operation.
func (o Operation) String() string {
	return fmt.Sprintf(
		"%s %s %s",
		o.leftSide.name,
		o.operator,
		o.rightSide.name,
	)
}

// newOperation creates a new operation.
func newOperation(operator Operator, leftSide, rightSide *Monkey) *Operation {
	return &Operation{
		operator:  operator,
		leftSide:  leftSide,
		rightSide: rightSide,
	}
}

// Monkey represents a monkey.
type Monkey struct {
	name      string
	job       Job
	number    int
	operation *Operation
	result    int
}

// String returns the string representation of a monkey.
func (m Monkey) String() string {
	if m.job == YellNumber {
		return fmt.Sprintf("%s: %d\n", m.name, m.number)
	}

	return fmt.Sprintf(
		"%s: %s\n",
		m.name,
		m.operation,
	)
}

// newMonkey creates a new monkey.
func newMonkey(name string, job Job, number int, operation *Operation) *Monkey {
	return &Monkey{
		name:      name,
		job:       job,
		number:    number,
		operation: operation,
	}
}

type MonkeyList map[string]*Monkey

// getMonkey returns a monkey from a list.
func (ml MonkeyList) getMonkey(name string) *Monkey {
	if _, ok := ml[name]; !ok {
		ml[name] = newMonkey(name, YellNumber, 0, nil)
	}

	return ml[name]
}

// setMonkey sets a monkey in a list.
func (ml MonkeyList) setMonkey(monkey *Monkey) {
	if _, ok := ml[monkey.name]; ok {
		ml[monkey.name].job = monkey.job
		ml[monkey.name].number = monkey.number
		ml[monkey.name].operation = monkey.operation
		return
	}

	ml[monkey.name] = monkey
}

// parseOperation parses an operation.
func (ml MonkeyList) parseOperation(operation string) *Operation {
	var operator Operator
	var leftSide, rightSide *Monkey

	operationParts := strings.Split(operation, " ")
	leftSide = ml.getMonkey(operationParts[0])
	operator = newOperator(operationParts[1])
	rightSide = ml.getMonkey(operationParts[2])

	return newOperation(operator, leftSide, rightSide)
}

// getMonkeyFromLine returns a monkey from a line.
func (ml MonkeyList) getMonkeyFromLine(line string) {
	job := YellNumber
	var operation *Operation
	monkeyAndJob := strings.Split(line, ":")
	name := strings.Trim(monkeyAndJob[0], " ")
	jobString := strings.Trim(monkeyAndJob[1], " ")

	number, err := strconv.Atoi(jobString)
	if err != nil {
		job = YellOperation
		operation = ml.parseOperation(jobString)
	}

	ml.setMonkey(newMonkey(name, job, number, operation))
}

// getResultForMonkey solves for the result of a monkey.
func (ml MonkeyList) getResultForMonkey(name string) {
	monkey := ml.getMonkey(name)

	if monkey.result != 0 {
		return
	}

	if monkey.job == YellNumber {
		monkey.result = monkey.number
		return
	}

	ml.getResultForMonkey(monkey.operation.leftSide.name)
	ml.getResultForMonkey(monkey.operation.rightSide.name)
	monkey.result = monkey.operation.getResult()
}

// String returns the string representation of a list of monkeys.
func (ml MonkeyList) String() string {
	result := ""

	for _, monkey := range ml {
		result += monkey.String()
	}

	return result
}

// getMonkeysFromFile returns a list of monkeys from a file.
func getMonkeysFromFile(txtlines []string) *MonkeyList {
	monkeys := make(MonkeyList, len(txtlines))

	for _, line := range txtlines {
		monkeys.getMonkeyFromLine(line)
	}

	return &monkeys
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	monkeys := getMonkeysFromFile(txtlines)
	if VERBOSE {
		fmt.Println(monkeys)
	}

	// part 1
	monkeys.getResultForMonkey(ROOT)
	root := monkeys.getMonkey(ROOT)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		root.result,
	)
}
