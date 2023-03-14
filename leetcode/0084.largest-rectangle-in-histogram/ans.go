// Package leetcode
// @Author: SalHe
// @Date:   2023/3/13 11:22
package leetcode

// 已查看【参考答案】单调栈
func largestRectangleArea(heights []int) int {
	count := len(heights)
	if count == 0 {
		return 0
	}
	if count == 1 {
		return heights[0]
	}

	// 该问题归其本质，是希望找到heights[i]能往左left、右right扩展的范围
	// 那么对于heights[i]，有heights[left]、heights[right]均<=heights[i]，
	// 并且left、right之间宽度尽可能大
	// 换言之，那么需要找到的right应该是最靠左边满足heights[i]>height[right+1]的right
	// left则是最靠右边满足heights[i]>heights[left-1]的left
	stack := []int{0} // {栈底, .... , 栈顶}
	area := 0
	heights = append([]int{0}, heights...)
	heights = append(heights, 0)
	// {0, ...heights..., 0}
	for i := 1; i < len(heights); i++ {
		height := heights[i]

		// 这里是求的已经加入到栈中的i对应高度可以获得的最大面积
		// 栈内的矩形块不可能扩展到当前i（右边界确定）
		for heights[stack[len(stack)-1]] > height {
			h := heights[stack[len(stack)-1]]
			stack = stack[0 : len(stack)-1] // 出栈
			area = max(area, h*(i-stack[len(stack)-1]-1))
		}

		stack = append(stack, i)
	}

	return area
}

// 超出时间限制
// N/A	N/A	Go	2023/03/13 11:49
// 添加备注

func largestRectangleAreaDp(heights []int) int {
	count := len(heights)
	if count <= 0 {
		return 0
	}
	dp := make([]int, count)
	maxArea := heights[0]
	diff := false
	for i := 0; i < count; i++ {
		dp[i] = heights[i]
		if maxArea != heights[i] {
			diff = true
		}
		maxArea = max(maxArea, heights[i])
	}
	if !diff {
		return count * maxArea
	}

	dp[0] = heights[0]
	// for i := 1; i < count; i++ {
	// 	dp[0][i] = min(dp[0][i-1], heights[i])
	// 	dp[i][i] = heights[i]
	// }

	for step := 2; step <= count; step++ {
		for start := 0; start < count+1-step; start++ {
			end := start + step - 1
			dp[start] = min(dp[start], heights[end])
			maxArea = max(maxArea, step*dp[start])
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
