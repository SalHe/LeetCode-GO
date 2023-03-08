package leetcode

func generateMatrix(n int) [][]int {
	result := make([][]int, n)
	for i := 0; i < len(result); i++ {
		result[i] = make([]int, n)
	}

	x, y := 0, 0
	dx, dy := 1, 0

	expectedSteps := n
	shouldDown := true
	restSteps := expectedSteps

	for i := 0; i < n*n; i++ {
		result[y][x] = i + 1

		restSteps--
		if restSteps == 0 {
			if shouldDown {
				expectedSteps--
			}
			shouldDown = !shouldDown
			restSteps = expectedSteps

			if dx == 1 && dy == 0 {
				dx, dy = 0, 1
			} else if dx == 0 && dy == 1 {
				dx, dy = -1, 0
			} else if dx == -1 && dy == 0 {
				dx, dy = 0, -1
			} else if dx == 0 && dy == -1 {
				dx, dy = 1, 0
			}
		}
		x += dx
		y += dy
	}

	return result
}
