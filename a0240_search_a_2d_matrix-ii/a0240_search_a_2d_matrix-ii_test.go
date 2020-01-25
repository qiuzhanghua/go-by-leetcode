package a0240_search_a_2d_matrix_ii

import "testing"

func TestSearchMatrix(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	if !searchMatrix(matrix, 5) {
		t.Errorf("Test failed, expected: %v, got:  '%v'", true, false)
	}
}

func TestSearchMatrix_2(t *testing.T) {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}

	if searchMatrix(matrix, 20) {
		t.Errorf("Test failed, expected: %v, got:  '%v'", false, true)
	}

}
