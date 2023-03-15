// Package leetcode
// @Author: SalHe
// @Date:   2023/3/15 10:36
package leetcode

type node struct {
	w, h int
}

// 提交结果	执行用时	内存消耗	语言	提交时间	备注
// 通过
// 8 ms	6 MB	Go	2023/03/15 11:14
// 添加备注
func maximalRectangle(matrix [][]byte) int {
	rows := len(matrix)
	cols := len(matrix[0])
	dp := make([][]node, rows+1)
	for i := 0; i < rows+1; i++ {
		dp[i] = make([]node, cols+1)
	}

	maxArea := 0
	for y, bytes := range matrix {
		for x, b := range bytes {
			nd := &dp[y+1][x+1]
			if b == '1' {
				nd.w = dp[y+1][x].w + 1
				nd.h = dp[y][x+1].h + 1
			}

			w, h := nd.w, nd.h
			if w*h <= maxArea {
				continue
			}

			for i := 0; i < w; i++ {
				h = min(dp[y+1][x+1-i].h, h)
				maxArea = max(maxArea, h*(i+1))
			}
		}
	}

	return maxArea
}

func min(a ...int) int {
	r := a[0]
	for _, i := range a {
		if i < r {
			r = i
		}
	}
	return r
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
