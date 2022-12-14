package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"

	"github.com/joaocarmo/advent-of-code/helpers"
)

const VERVOSE = false
const ERR_UNDETERMINED = "Undetermined"
const EXTRA_LINES = "[[2]]\n[[6]]\n"
const NEW_LINE = "\n"

// Pair is a pair of packets.
type Pair struct {
	Left       []interface{}
	Right      []interface{}
	AreInOrder bool
}

// validateOrder validates if the pair is in order.
func (p *Pair) validateOrder() {
	p.AreInOrder, _ = comparePackets(p.Left, p.Right)

	if VERVOSE {
		fmt.Println("------------------------------------------------", p.AreInOrder)
	}
}

// String returns a string representation of the pair.
func (p *Pair) String() string {
	return fmt.Sprintf("%v\t%v\t%t", p.Left, p.Right, p.AreInOrder)
}

// newPair creates a new pair.
func newPair(left, right []interface{}) *Pair {
	return &Pair{
		Left:       left,
		Right:      right,
		AreInOrder: false,
	}
}

// comparePackets compares two packets.
func comparePackets(left, right interface{}) (bool, error) {
	if VERVOSE {
		fmt.Printf("--> %v vs. %v\n", left, right)
	}

	leftSlice, _ := left.([]interface{})
	rightSlice, _ := right.([]interface{})

	for i := 0; i < helpers.MaxOf(len(leftSlice), len(rightSlice)); i++ {
		if i == len(leftSlice) {
			return true, nil
		}

		if i == len(rightSlice) {
			return false, nil
		}

		newLeft := leftSlice[i]
		newRight := rightSlice[i]

		leftFloat, leftIsFloat := newLeft.(float64)
		rightFloat, rightIsFloat := newRight.(float64)
		newLeftSlice, _ := newLeft.([]interface{})
		newRightSlice, _ := newRight.([]interface{})

		if leftIsFloat && rightIsFloat {
			leftInt := int(leftFloat)
			rightInt := int(rightFloat)

			if leftInt != rightInt {
				return leftInt < rightInt, nil
			}
		} else {
			if leftIsFloat {
				newLeftSlice = append(newLeftSlice, leftFloat)
			}

			if rightIsFloat {
				newRightSlice = append(newRightSlice, rightFloat)
			}

			result, err := comparePackets(newLeftSlice, newRightSlice)

			if err == nil {
				return result, nil
			}
		}
	}

	return false, fmt.Errorf(ERR_UNDETERMINED)
}

// parsePacketFromLine parses a packet from a line.
func parsePacketFromLine(line string) []interface{} {
    var packet []interface{}

    err := json.Unmarshal([]byte(line), &packet)

    if err != nil {
		fmt.Println(err)
	}

	return packet
}

// getPacketFromLines gets a packet from two lines.
func getPacketFromLines(lines []string) ([]interface{}, []interface{}) {
	left := parsePacketFromLine(lines[0])
	right := parsePacketFromLine(lines[1])

	return left, right
}

// getPairsFromFile gets all pairs from a file.
func getPairsFromFile(txtlines []string) []*Pair {
	pairs := []*Pair{}

	var currentPair []string
	for _, line := range txtlines {
		if line == "" {
			continue
		}

		if len(currentPair) == 2 {
			left, right := getPacketFromLines(currentPair)
			pairs = append(pairs, newPair(left, right))
			currentPair = []string{}
		}

		currentPair = append(currentPair, line)
	}

	left, right := getPacketFromLines(currentPair)
	pairs = append(pairs, newPair(left, right))

	return pairs
}

// getPacketsFromFile gets all packets from a file.
func getPacketsFromFile(txtlines []string) []interface{} {
	packets := []interface{}{}

	for _, line := range txtlines {
		if line == "" {
			continue
		}

		packet := parsePacketFromLine(line)
		packets = append(packets, packet)
	}

	return packets
}

// getIndicesInRightOrder gets the indices of the pairs in the right order.
func getIndicesInRightOrder(pairs []*Pair) []int {
	indices := []int{}

	for i, pair := range pairs {
		index := i + 1

		if VERVOSE {
			fmt.Printf("~ Index %d\n", index)
		}

		pair.validateOrder()

		if pair.AreInOrder {
			indices = append(indices, index)
		}
	}

	return indices
}

// sortPairsByRightOrder sorts the pairs by the right order.
func sortPairsByRightOrder(packets []interface{}) []interface{} {
	sort.SliceStable(packets, func(i, j int) bool {
		result, _ := comparePackets(packets[i], packets[j])
		return result
	})

	return packets
}

// findDividerIndices finds the indices of the dividers.
func findDividerIndices(packets, dividers []interface{}) []int {
	indices := []int{}

	for i, packet := range packets {
		index := i + 1
		for _, divider := range dividers {
			if reflect.DeepEqual(packet, divider) {
				indices = append(indices, index)
			}
		}
	}

	return indices
}

// main is the entry point for the application.
func main() {
	// read the file
	args := helpers.ReadArguments()
	filename := args[0]
	txtlines := helpers.ReadFile(filename)

	// process the file
	pairs := getPairsFromFile(txtlines)

	// part 1
	indicesInRightOrder := getIndicesInRightOrder(pairs)
	sumOfIndicesInRightOrder := helpers.SumInts(indicesInRightOrder...)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		sumOfIndicesInRightOrder,
	)

	// process the file again
	extraLines := strings.Split(EXTRA_LINES, NEW_LINE)
	fullTxtlines := append(txtlines, extraLines...)

	// part 2
	packets := getPacketsFromFile(fullTxtlines)
	packetsInRightOrder := sortPairsByRightOrder(packets)
	dividers := getPacketsFromFile(extraLines)
	dividerIndices := findDividerIndices(packetsInRightOrder, dividers)
	decoderKey := helpers.MultiplyInts(dividerIndices...)
	fmt.Printf(
		"[Part One] The answer is: %d\n",
		decoderKey,
	)
}
