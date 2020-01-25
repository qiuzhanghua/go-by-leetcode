package a0240_search_a_2d_matrix_ii

func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) < 1 || len(matrix[0]) < 1 {
		return false
	}
	col := len(matrix[0]) - 1
	row := 0

	for col >= 0 && row <= len(matrix)-1 {
		if matrix[row][col] < target {
			row++
		} else if matrix[row][col] > target {
			col--
		} else {
			return true
		}
	}

	return false
}
