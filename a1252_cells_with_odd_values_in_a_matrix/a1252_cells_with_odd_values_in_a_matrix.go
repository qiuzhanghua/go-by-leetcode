package a1252_cells_with_odd_values_in_a_matrix

func oddCells(n int, m int, indices [][]int) int {
	row, col := make([]bool, n), make([]bool, m)
	r, c := 0, 0
	for _, p := range indices {
		x, y := p[0], p[1]
		row[x] = row[x] != true
		col[y] = col[y] != true
		if row[x] {
			r += 1
		} else {
			r += -1
		}
		if col[y] {
			c += 1
		} else {
			c += -1
		}
	}
	return r*m + c*n - 2*r*c
}
