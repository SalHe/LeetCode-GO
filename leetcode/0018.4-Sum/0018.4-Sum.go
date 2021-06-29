package leetcode

import "sort"

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	sort.Ints(nums)
	n := len(nums)

	for h := 0; h < n; h++ {
		if h > 0 && nums[h] == nums[h-1] {
			continue
		}

		for i := h + 1; i < n; i++ {

			if i > h+1 && nums[i] == nums[i-1] {
				continue
			}

			nowTarget := target - (nums[i] + nums[h])
			k := n - 1
			for j := i + 1; j < n; j++ {

				if j > i+1 && nums[j] == nums[j-1] {
					continue
				}

				for nums[j]+nums[k] > nowTarget && k > j {
					k--
				}

				if k == j {
					break
				}
				if nums[j]+nums[k] == nowTarget {
					ans = append(ans, []int{nums[h], nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return ans
}
