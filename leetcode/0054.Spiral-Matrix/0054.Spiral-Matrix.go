package leetcode

func spiralOrder(matrix [][]int) (ans []int) {
	m := len(matrix)
	if m <= 0 {
		return
	}
	n := len(matrix[0])

	ans = make([]int, m*n)
	dx, dy := 1, 0 // go right
	x, y := 0, 0
	xBoard := []int{0, n - 1}
	yBoard := []int{1, m - 1}
	for i := 0; i < m*n; i++ {
		ans[i] = matrix[y][x]
		if dx == 1 && dy == 0 && x >= xBoard[1] {
			x = xBoard[1]
			xBoard[1]--
			dx, dy = 0, 1
		} else if dx == 0 && dy == 1 && y >= yBoard[1] {
			y = yBoard[1]
			yBoard[1]--
			dx, dy = -1, 0
		} else if dx == -1 && dy == 0 && x <= xBoard[0] {
			x = xBoard[0]
			xBoard[0]++
			dx, dy = 0, -1
		} else if dx == 0 && dy == -1 && y <= yBoard[0] {
			y = yBoard[0]
			yBoard[0]++
			dx, dy = 1, 0
		}
		x += dx
		y += dy
	}

	return
}
