package a0756_pyramid_transition_matrix

const AlphabeticSize = 7

func pyramidTransition(bottom string, allowed []string) bool {
	tr := make([][]rune, AlphabeticSize)
	for i := range tr {
		tr[i] = make([]rune, AlphabeticSize)
	}
	for _, s := range allowed {
		tr[s[0]-'A'][s[1]-'A'] |= 1 << (s[2] - 'A')
	}
	return pyramid([]rune(bottom), tr, 0, make([]rune, len(bottom)-1))
}

func pyramid(row []rune, tr [][]rune, i int, nextRow []rune) bool {
	if len(row) == 1 {
		return true
	}
	if i == len(row)-1 {
		return pyramid(nextRow, tr, 0, make([]rune, len(nextRow)-1))
	}
	bit := tr[row[i]-'A'][row[i+1]-'A']
	for j := 0; j < AlphabeticSize; j++ {
		if (bit & (1 << uint(j))) != 0 {
			nextRow[i] = rune('A' + j)
			if pyramid(row, tr, i+1, nextRow) {
				return true
			}
		}
	}

	return false
}
