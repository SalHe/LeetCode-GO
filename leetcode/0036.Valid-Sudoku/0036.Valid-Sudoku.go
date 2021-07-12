package leetcode

func isValidSudoku(board [][]byte) bool {
	row := make(map[byte]int)
	columns := make([]map[byte]int, 9)
	boxes := make([]map[byte]int, 9)
	for i := 0; i < 9; i++ {
		columns[i] = make(map[byte]int, 9)
		boxes[i] = make(map[byte]int, 9)
	}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			num := board[i][j]
			if num == '.' {
				continue
			}
			box := 3*(i/3) + j/3
			if r, ok := row[num]; ok && r >= i {
				return false
			}
			if _, ok := columns[j][num]; ok {
				return false
			}
			if _, ok := boxes[box][num]; ok {
				return false
			}
			row[num] = i
			// 下面两个无所谓
			columns[j][num] = 0
			boxes[box][num] = 0
		}
	}
	return true
}
