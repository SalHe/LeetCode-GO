package leetcode

func insert(intervals [][]int, newInterval []int) (result [][]int) {
	count := len(intervals)
	if count <= 0 {
		result = append(result, newInterval)
		return
	}

	i := 0
	for i < count && intervals[i][1] < newInterval[0] {
		result = append(result, intervals[i])
		i++
	}

	// newInterval := []int{0, 0}
	for i < count && intervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], intervals[i][0])
		newInterval[1] = max(newInterval[1], intervals[i][1])
		i++
	}

	result = append(result, newInterval)
	result = append(result, intervals[i:]...)

	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
