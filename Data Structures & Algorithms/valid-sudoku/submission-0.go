func isValidSudoku(board [][]byte) bool {
	var rows, cols, boxes [9]int

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c ++ {
			// ignore the empty cell
			val := board[r][c]
			if  val == '.' {
				continue
			}
			
			mask := 1 << (val - '0')

			boxIdx := (r / 3) * 3 + (c / 3)

			// check duplicate
			if (rows[r] & mask) != 0 ||
				(cols[c] & mask) != 0 ||
				(boxes[boxIdx] & mask) != 0 {
				return false
			}

			rows[r] |= mask
			cols[c] |= mask
			boxes[boxIdx] |= mask
		}
	}
	return true
}
