package leetcode

// 这里主要考虑思路的正确性，就不使用位优化了
func solveSudoku(board [][]byte) {
	var columns [9][9]bool
	var rows [9][9]bool
	var boxes [3][3][9]bool
	var spaces [][2]int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			ch := board[i][j]
			if ch != '.' {
				num := ch - '0'
				rows[i][num-1] = true
				columns[j][num-1] = true
				boxes[i/3][j/3][num-1] = true
			} else {
				spaces = append(spaces, [2]int{i, j})
			}
		}
	}

	var dfs func(int) bool
	dfs = func(iSpace int) bool {
		if iSpace == len(spaces) {
			return true
		}
		x, y := spaces[iSpace][1], spaces[iSpace][0]
		for i := 0; i < 9; i++ {
			if !columns[x][i] && !rows[y][i] && !boxes[y/3][x/3][i] {
				rows[y][i] = true
				columns[x][i] = true
				boxes[y/3][x/3][i] = true
				board[y][x] = byte('1' + i)
				if dfs(iSpace + 1) {
					return true
				}
				rows[y][i] = false
				columns[x][i] = false
				boxes[y/3][x/3][i] = false
			}
		}
		return false
	}

	dfs(0)
}
