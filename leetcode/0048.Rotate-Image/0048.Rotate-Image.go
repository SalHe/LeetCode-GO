package leetcode

func rotate(matrix [][]int) {
	n := len(matrix)

	// 转置
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 对称
	var a int
	if n%2 == 0 {
		a = n / 2
	} else {
		a = (n + 1) / 2
	}
	for i := 0; i < n; i++ {
		for j := 0; j < a; j++ {
			another := n - 1 - j
			matrix[i][j], matrix[i][another] = matrix[i][another], matrix[i][j]
		}
	}
}
