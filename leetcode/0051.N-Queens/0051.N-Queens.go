package leetcode

func solveNQueens(n int) (solutions [][]string) {
	if n <= 0 {
		return
	}
	placements := make([]int, n)
	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == n {
			solution := make([]string, n)
			line := make([]rune, n)
			for row := 0; row < n; row++ {
				for p := 0; p < n; p++ {
					if placements[row] == p {
						line[p] = 'Q'
					} else {
						line[p] = '.'
					}
				}
				solution[row] = string(line)
			}
			solutions = append(solutions, solution)
			return
		}

		for position := 0; position < n; position++ {
			for row := 0; row < depth; row++ {
				// 检查选择的位置（position）会不会与之前已放置的皇后冲突
				if placements[row] == position || placements[row] == position-(depth-row) || placements[row] == position+(depth-row) {
					goto nextPosition
				}
			}
			placements[depth] = position
			dfs(depth + 1)

		nextPosition:
		}
	}
	for i := 0; i < n; i++ {
		placements[0] = i
		dfs(1)
	}
	return
}
