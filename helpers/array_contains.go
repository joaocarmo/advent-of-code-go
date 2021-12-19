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

// StrArrayContains checks whether an string belogs to an array of string.
func StrArrayContains(arr []string, el string) bool {
	for _, arrEl := range arr {
		if arrEl == el {
			return true
		}
	}

	return false
}
