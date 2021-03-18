package hello

func LongestSubStrWithoutRepeatingChars(s string) int {
	lastOccurred := make(map[rune]int)
	maxLength := 0
	start := 0
	for index, value := range []rune(s) {

		if lastI, ok := lastOccurred[value]; ok && lastI >= start {
			start = lastOccurred[value] + 1
		}

		if index-start+1 > maxLength {
			maxLength = index - start + 1
		}

		lastOccurred[value] = index
	}
	return maxLength
}
