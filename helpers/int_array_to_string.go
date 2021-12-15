package helpers

import (
	"strconv"
	"strings"
)

// IntArrayToString converts a string to an array of ints.
func IntArrayToString(intArray []int, separator string) string {
	var strArray []string

	for _, num := range intArray {
		strArray = append(strArray, strconv.Itoa(num))
	}

	return strings.Join(strArray, separator)
}
