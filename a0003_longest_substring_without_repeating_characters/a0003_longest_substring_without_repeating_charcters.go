package a0003_longest_substring_without_repeating_characters

func lengthOfLongestSubstring(s string) int {
	// key is rune and value is index
	letters := make(map[rune]int)
	// convert string to slice of runes
	runes := []rune(s)
	n, result := len(s), 0

	for i, j := 0, 0; j < n; j++ {
		r := runes[j]
		if pos, ok := letters[r]; ok {
			// only move i forward
			if pos >= i {
				i = pos + 1
			}
		}

		if l := j - i + 1; l > result {
			result = l
		}

		letters[r] = j
	}

	return result
}
