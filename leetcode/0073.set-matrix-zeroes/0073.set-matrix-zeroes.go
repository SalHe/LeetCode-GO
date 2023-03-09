package leetcode

// 优化思路：使用第一列和第一行来标记（但是可能要额外设计一下算法）

func setZeroes(matrix [][]int) {
	rows := make([]bool, len(matrix))
	cols := make([]bool, len(matrix[0]))
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[y][x] == 0 {
				rows[y] = true
				cols[x] = true
			}
		}
	}

	for y := 0; y < len(rows); y++ {
		if !rows[y] {
			continue
		}
		for x := 0; x < len(cols); x++ {
			matrix[y][x] = 0
		}
	}

	for x := 0; x < len(cols); x++ {
		if !cols[x] {
			continue
		}
		for y := 0; y < len(rows); y++ {
			matrix[y][x] = 0
		}
	}
}
