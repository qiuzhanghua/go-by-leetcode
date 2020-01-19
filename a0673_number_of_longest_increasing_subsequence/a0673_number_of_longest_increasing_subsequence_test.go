package a0673_number_of_longest_increasing_subsequence

import "testing"

func TestFindNumberOfLIS(t *testing.T) {
	expected := 2
	actual := findNumberOfLIS([]int{1, 3, 5, 4, 7})
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}

func TestFindNumberOfLIS_2(t *testing.T) {
	expected := 5
	actual := findNumberOfLIS([]int{2, 2, 2, 2, 2})
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
