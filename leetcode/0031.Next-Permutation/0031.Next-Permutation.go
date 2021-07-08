package leetcode

import "sort"

func nextPermutation(nums []int) {
	if len(nums) <= 1 {
		return
	}

	n := len(nums)
	i, j, k := n-2, n-1, n-1
	for i >= 0 && nums[i] >= nums[j] {
		i--
		j--
	}

	if i >= 0 {
		for nums[i] >= nums[k] {
			k--
		}
		nums[i], nums[k] = nums[k], nums[i]
	}

	for i, j := j, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}

}

func nextPermutation_foolish(nums []int) {
	// sorts := make([]int, 101) // 0 <= nums[i] <= 100
	// for _, num := range nums {
	// 	sorts[num]++
	// }
	// for i := 1; i <= 100; i++ {
	// 	sorts[i] += sorts[i-1]
	// }
	n := len(nums)
	radix := make([]int, n)
	copy(radix, nums)
	sort.Ints(radix)
	indexes := make([]int, n)
	{
		used := make([]bool, n)
		for i := n - 1; i >= 0; i-- {
			for j := 0; j < n; j++ {
				if !used[j] && radix[j] == nums[i] {
					indexes[i] = j
					used[j] = true
					break
				}
			}
		}

	}

	for {
		increaseRadix(indexes, n, n-1)
		table := make(map[int]int)
		pass := true
		for _, index := range indexes {
			if _, ok := table[index]; ok {
				pass = false
				break
			}
			table[index] = 0
		}

		if pass {
			break
		}
	}

	for i := range nums {
		nums[i] = radix[indexes[i]]
	}

}

func increaseRadix(indexes []int, max int, i int) {
	if i < 0 {
		return
	}
	indexes[i]++
	if indexes[i] >= max {
		indexes[i] = 0
		increaseRadix(indexes, max, i-1)
	}
}
