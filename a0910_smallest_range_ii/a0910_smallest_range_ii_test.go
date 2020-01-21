package a0910_smallest_range_ii

import "testing"

func TestSmallestRangeII(t *testing.T) {
	expected := 0
	actual := smallestRangeII([]int{1}, 0)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got:  '%v'", expected, actual)
	}
}

func TestSmallestRangeII_2(t *testing.T) {
	expected := 6
	actual := smallestRangeII([]int{0, 10}, 2)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got:  '%v'", expected, actual)
	}
}

func TestSmallestRangeII_3(t *testing.T) {
	expected := 3
	actual := smallestRangeII([]int{1, 3, 6}, 3)
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got:  '%v'", expected, actual)
	}
}
