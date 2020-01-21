package a0999_available_captures_for_rook

import "testing"

func TestNumRookCaptures(t *testing.T) {
	board := [][]byte{
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', 'R', '.', '.', '.', 'p'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', 'p', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.'},
	}
	expected := 3

	actual := numRookCaptures(board)
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
