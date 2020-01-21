package a1252_cells_with_odd_values_in_a_matrix

import "testing"

func TestOddCells(t *testing.T) {
	expected := 6
	actual := oddCells(2, 3, [][]int{
		{0, 1},
		{1, 1},
	})
	if actual != expected {
		t.Errorf("Test failed, expected: '%v', got:  '%v'", expected, actual)
	}
}
