package helpers

// IntArrayContains checks whether an integer belogs to an array of integers.
func IntArrayContains(arr []int, el int) bool {
	for _, arrEl := range arr {
		if arrEl == el {
			return true
		}
	}

	return false
}

// StrArrayContains checks whether a string belogs to an array of strings.
func StrArrayContains(arr []string, el string) bool {
	for _, arrEl := range arr {
		if arrEl == el {
			return true
		}
	}

	return false
}

// RuneArrayContains checks whether a rune belogs to an array of runes.
func RuneArrayContains(arr []rune, el rune) bool {
	for _, arrEl := range arr {
		if arrEl == el {
			return true
		}
	}

	return false
}
