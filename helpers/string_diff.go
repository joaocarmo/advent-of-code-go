package helpers

import "strings"

// StringDiff returns the character difference between two strings.
func StringDiff(str1 string, str2 string) []string {
	var diff []string

	str1Arr := strings.Split(str1, "")
	str2Arr := strings.Split(str2, "")

	for _, c := range str1Arr {
		if !StrArrayContains(str2Arr, c) && !StrArrayContains(diff, c) {
			diff = append(diff, c)
		}
	}

	for _, c := range str2Arr {
		if !StrArrayContains(str1Arr, c) && !StrArrayContains(diff, c) {
			diff = append(diff, c)
		}
	}

	return diff
}
