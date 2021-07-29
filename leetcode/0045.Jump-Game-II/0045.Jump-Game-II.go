package leetcode

import "math"

func jump(nums []int) int {
	n := len(nums)
	times := make([]int, n)
	last := n - 1
	couldReach := func(curr, target int) bool {
		return target-curr <= nums[curr]
	}
	for i := last - 1; i >= 0; i-- {
		times[i] = math.MaxInt64
		if couldReach(i, last) {
			times[i] = 1
		} else {
			for j := last - 1; j > i; j-- {
				// 尝试寻找路线：i -> j -> last
				if couldReach(i, j) && times[j] < math.MaxInt64 {
					times[i] = min(times[j]+1, times[i])
				}
			}
		}
	}
	times[last] = 0

	return times[0]
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
