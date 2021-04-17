package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)

	for i := 0; i < n; i++ {

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		k := n - 1
		for j := i + 1; j < n; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for nums[j]+nums[k] > target && k > j {
				k--
			}

			if k == j {
				break
			}
			if nums[j]+nums[k] == target {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
			}
		}
	}
	return ans
}
