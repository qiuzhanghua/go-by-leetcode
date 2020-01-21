package a0999_available_captures_for_rook

func numRookCaptures(board [][]byte) int {
	count := 0
	i := 0
	j := 0
	for a, row := range board {
		for b, cell := range row {
			if string(cell) == "R" {
				i = a
				j = b
				break
			}
		}
	}
	for _, directions := range [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		x := i + directions[0]
		y := j + directions[1]
		for x >= 0 && y >= 0 && x < len(board) && y < len(board[0]) && (board[x][y] == '.' || board[x][y] == 'p') {
			if board[x][y] == 'p' {
				count += 1
				break
			}
			x += directions[0]
			y += directions[1]
		}
	}
	return count
}
