package helpers

import (
	"strconv"
	"strings"
)

// StringToIntArray converts a string to an array of ints.
func StringToIntArray(str string, separator string) []int {
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
