// Package leetcode
// @Author: SalHe
// @Date:   2023/3/15 10:36
package leetcode

// 通过
// 4 ms	6.1 MB	Go	2023/03/15 11:38
// DP + 最大矩形面积
func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]int, rows+1)
	for i := 0; i < rows+1; i++ {
		dp[i] = make([]int, cols+1+1)
	}

	maxArea := 0
	for y, bytes := range matrix {
		stack := []int{0}
		bytes = append(bytes, '0')
		for x, b := range bytes {
			if b == '1' {
				dp[y+1][x+1] = dp[y][x+1] + 1
			}

			h := dp[y+1][x+1]
			for dp[y+1][stack[len(stack)-1]] > h {
				height := dp[y+1][stack[len(stack)-1]]
				stack = stack[:len(stack)-1]
				maxArea = max(maxArea, height*(x-stack[len(stack)-1]))
			}
			stack = append(stack, x+1)
		}
	}

	return maxArea
}

func max(a ...int) int {
	r := a[0]
	for _, i := range a {
		if i > r {
			r = i
		}
	}
	return r
}
