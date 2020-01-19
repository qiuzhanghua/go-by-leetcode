package a0279_perfect_squares

import "testing"

func TestPerfectSquares(t *testing.T) {
	expected := 3
	actual := numSquares(12)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}

	expected = 2
	actual = numSquares(13)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
