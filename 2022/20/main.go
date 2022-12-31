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
const DECRYPTION_KEY_1 = 1
const DECRYPTION_KEY_2 = 811589153
const MOVE_TIMES_1 = 1
const MOVE_TIMES_2 = 10

type Number struct {
	value int
}

func (n Number) String() string {
	return strconv.Itoa(n.value)
}

// newNumber creates a new Number.
func newNumber(value int) *Number {
	return &Number{value: value}
}

type Message struct {
	encrypted []*Number
	decrypted []*Number
}

// getNextIndex returns the next index of the number in the decrypted message.
func (m Message) getNextIndex(n *Number, currentIndex int) int {
	numElements := len(m.encrypted)
	lastIndex := numElements - 1

	return helpers.EuclideanRemainder(currentIndex + n.value, lastIndex)
}

// getGroveCoordinate returns the grove coordinate of the decrypted message.
func (m Message) getGroveCoordinate(n int) *Number {
	zeroIndex := findIndexByValue(m.decrypted, 0)
	decryptedIndex := (zeroIndex + n) % len(m.decrypted)

	if VERBOSE {
		fmt.Println("--> zeroIndex:", zeroIndex)
		fmt.Println("--> decryptedIndex:", decryptedIndex)
	}

	return m.decrypted[decryptedIndex]
}

// getSumOfCoordinates returns the sum of the grove coordinates.
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

// getRealIndex returns the real index of the number in the encrypted message.
func (m Message) getRealIndex(n int) int {
	return n % len(m.encrypted)
}

// getRealValue returns the real value of the number in the encrypted message.
func (m Message) getRealValue(n int) *Number {
	return m.encrypted[m.getRealIndex(n)]
}

// moveIndex moves the index of the number in the decrypted message.
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

// moveAllIndexes moves all indexes of the numbers in the decrypted message.
func (m Message) moveAllIndexes(times int) {
	for j := 0; j < times; j++ {
		for i := 0; i < len(m.encrypted); i++ {
			m.moveIndex(i)

			if VERY_VERBOSE {
				fmt.Println("------------------------------------- Moving index:", i, "(#",  j, ")")
				fmt.Println(m)
			}
		}
	}
}

// String returns the string representation of the message.
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

// newMessage creates a new Message.
func newMessage(encrypted []*Number, key int) Message {
	encryptedWithKey := multiplyAllByNumber(encrypted, key)
	decryptedWithKey := make([]*Number, len(encryptedWithKey))
	copy(decryptedWithKey, encryptedWithKey)
	m := Message{
		encrypted: encryptedWithKey,
		decrypted: decryptedWithKey,
	}
	return m
}

// multiplyAllByNumber multiplies all numbers in the array by a number.
func multiplyAllByNumber(numbers []*Number, n int) []*Number {
	for i, number := range numbers {
		numbers[i] = newNumber(number.value * n)
	}

	return numbers
}

// findIndexByValue returns the index of the number in the array.
func findIndexByValue(array []*Number, n int) int {
	for i, number := range array {
		if number.value == n {
			return i
		}
	}

	return -1
}

// findIndex returns the index of the number in the array.
func findIndex(array []*Number, n *Number) int {
	for i, number := range array {
		if number == n {
			return i
		}
	}

	return -1
}

// insert inserts a number in the array.
func insert(array []*Number, value *Number, index int) []*Number {
    return append(array[:index], append([]*Number{value}, array[index:]...)...)
}

// remove removes a number from the array.
func remove(array []*Number, index int) []*Number {
    return append(array[:index], array[index+1:]...)
}

// move moves a number in the array.
func move(array []*Number, srcIndex int, dstIndex int) []*Number {
    value := array[srcIndex]
    return insert(remove(array, srcIndex), value, dstIndex)
}

// getEncryptedMessageFromFile returns the encrypted message from the file.
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
	message1 := newMessage(encrypted, DECRYPTION_KEY_1)
	message1.moveAllIndexes(MOVE_TIMES_1)
	sumCoordinates1 := message1.getSumOfCoordinates(groveCoordinateIndices)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		sumCoordinates1,
	)

	// part 2
	message2 := newMessage(encrypted, DECRYPTION_KEY_2)
	message2.moveAllIndexes(MOVE_TIMES_2)
	sumCoordinates2 := message2.getSumOfCoordinates(groveCoordinateIndices)
	fmt.Printf(
		"[Part Two] The answer is: %d\n",
		sumCoordinates2,
	)
}
