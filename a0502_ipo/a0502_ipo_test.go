package a0502_ipo

import "testing"

func TestFindMaximizedCapital(t *testing.T) {
	expected := 4
	actual := findMaximizedCapital(2, 0, []int{1, 2, 3}, []int{0, 1, 1})
	if expected != actual {
		t.Errorf("Test failed, expected: %v, got:  '%v'", expected, actual)
	}

}
