package leetcode

func minPathSum(grid [][]int) int {
	// 求格子形状
	m := len(grid)    // >=1
	n := len(grid[0]) // >=1

	for y := 1; y < m; y++ {
		grid[y][0] += grid[y-1][0]
	}
	for x := 1; x < n; x++ {
		grid[0][x] += grid[0][x-1]
	}

	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if grid[y][x-1] < grid[y-1][x] {
				grid[y][x] += grid[y][x-1]
			} else {
				grid[y][x] += grid[y-1][x]
			}
		}
	}

	return grid[m-1][n-1]
}
