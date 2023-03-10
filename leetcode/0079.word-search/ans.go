package leetcode

func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	visited := make([][]bool, rows)
	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			if word[0] == board[y][x] && solve(board, rows, cols, []byte(word)[1:], visited, x, y) {
				return true
			}
		}
	}

	return false
}

var dirs = [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}

func solve(board [][]byte, rows int, cols int, toFind []byte, visited [][]bool, x int, y int) bool {
	if len(toFind) <= 0 {
		return true
	}

	visited[y][x] = true
	for _, dir := range dirs {
		yy := y + dir[1]
		xx := x + dir[0]
		if yy >= 0 && yy < rows &&
			xx >= 0 && xx < cols &&
			!(visited[yy][xx]) &&
			board[yy][xx] == toFind[0] &&
			solve(board, rows, cols, toFind[1:], visited, xx, yy) {
			return true
		}
	}
	visited[y][x] = false

	return false
}
