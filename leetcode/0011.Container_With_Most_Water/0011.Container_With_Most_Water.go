package leetcode

func maxArea(height []int) int {
	i, j := 0, len(height)-1
	area := 0
	for i <= j {
		if height[i] < height[j] {
			area = max(area, height[i]*(j-i))
			i++
		} else {
			area = max(area, height[j]*(j-i))
			j--
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
