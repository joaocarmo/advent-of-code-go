package main

import (
	"fmt"
	"strconv"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERBOSE = false
const VERY_VERBOSE = false
const GROVE_COORDINATE_1 = 1000
const GROVE_COORDINATE_2 = 2000
const GROVE_COORDINATE_3 = 3000

type Number struct {
	value int
}

func (n Number) String() string {
	return strconv.Itoa(n.value)
}

func newNumber(value int) *Number {
	return &Number{value: value}
}

type Message struct {
	encrypted []*Number
	decrypted []*Number
}

func (m Message) getNextIndex(n *Number, currentIndex int) int {
	numElements := len(m.encrypted)
	lastIndex := numElements - 1

	return helpers.EuclideanRemainder(currentIndex + n.value, lastIndex)
}

func (m Message) getGroveCoordinate(n int) *Number {
	zeroIndex := findIndexByValue(m.decrypted, 0)
	decryptedIndex := (zeroIndex + n) % len(m.decrypted)

	if VERBOSE {
		fmt.Println("--> zeroIndex:", zeroIndex)
		fmt.Println("--> decryptedIndex:", decryptedIndex)
	}

	return m.decrypted[decryptedIndex]
}

func (m Message) getSumOfCoordinates(indices []int) int {
	sumCoordinates := 0

	for _, groveCoordinateIndex := range indices {
		groveCoordinate := m.getGroveCoordinate(groveCoordinateIndex)
		sumCoordinates += groveCoordinate.value

		if VERBOSE {
			fmt.Println("Grove coordinate", groveCoordinateIndex, "is", groveCoordinate)
		}
	}

	return sumCoordinates
}

func (m Message) getRealIndex(n int) int {
	return n % len(m.encrypted)
}

func (m Message) getRealValue(n int) *Number {
	return m.encrypted[m.getRealIndex(n)]
}

func (m Message) moveIndex(n int) {
	number := m.getRealValue(n)

	indexFrom := findIndex(m.decrypted, number)
	indexTo := m.getNextIndex(number, indexFrom)

	if VERBOSE {
		fmt.Println("--> number:", number)
		fmt.Println("--> indexFrom:", indexFrom)
		fmt.Println("--> indexTo:", indexTo)
	}

	move(m.decrypted, indexFrom, indexTo)
}

func (m Message) moveAllIndexes() {
	for i := 0; i < len(m.encrypted); i++ {
		m.moveIndex(i)

		if VERY_VERBOSE {
			fmt.Println("------------------------------------- Moving index:", i)
			fmt.Println(m)
		}
	}
}

func (m Message) String() string {
	str := "Encrypted:\n"

	for _, number := range m.encrypted {
		str += number.String() + ", "
	}

	str += "\n\nDecrypted:\n"

	for _, number := range m.decrypted {
		str += number.String() + ", "
	}

	return str
}

func newMessage(encrypted []*Number) Message {
	decrypted := make([]*Number, len(encrypted))
	copy(decrypted, encrypted)
	m := Message{
		encrypted: encrypted,
		decrypted: decrypted,
	}
	return m
}

func findIndexByValue(array []*Number, n int) int {
	for i, number := range array {
		if number.value == n {
			return i
		}
	}

	return -1
}

func findIndex(array []*Number, n *Number) int {
	for i, number := range array {
		if number == n {
			return i
		}
	}

	return -1
}

func insert(array []*Number, value *Number, index int) []*Number {
    return append(array[:index], append([]*Number{value}, array[index:]...)...)
}

func remove(array []*Number, index int) []*Number {
    return append(array[:index], array[index+1:]...)
}

func move(array []*Number, srcIndex int, dstIndex int) []*Number {
    value := array[srcIndex]
    return insert(remove(array, srcIndex), value, dstIndex)
}

func getEncryptedMessageFromFile(txtlines []string) []*Number {
	encrypted := []*Number{}

	for _, line := range txtlines {
		numberInt, _ := strconv.Atoi(line)
		number := newNumber(numberInt)
		encrypted = append(encrypted, number)
	}

	return encrypted
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	encrypted := getEncryptedMessageFromFile(txtlines)

	// constants
	groveCoordinateIndices := []int{
		GROVE_COORDINATE_1,
		GROVE_COORDINATE_2,
		GROVE_COORDINATE_3,
	}

	// part 1
	message := newMessage(encrypted)
	message.moveAllIndexes()
	sumCoordinates := message.getSumOfCoordinates(groveCoordinateIndices)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		sumCoordinates,
	)
}
