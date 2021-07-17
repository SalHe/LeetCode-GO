package leetcode

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, num := range nums {
		if num <= 0 {
			nums[i] = n + 1
		}
	}

	for _, num := range nums {
		x := abs(num)
		if x <= n {
			nums[x-1] = -abs(nums[x-1])
		}
	}

	for i, num := range nums {
		if num > 0 {
			return i + 1
		}
	}
	return n + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
