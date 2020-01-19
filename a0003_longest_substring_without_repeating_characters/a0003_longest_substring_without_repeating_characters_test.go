package a0003_longest_substring_without_repeating_characters

import "testing"

func TestLongest(t *testing.T) {
	s := "abcabcbb"
	actual := lengthOfLongestSubstring(s)
	expected := 3
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
