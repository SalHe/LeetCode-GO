package leetcode

func removeDuplicates(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}
	i, n := 0, 1
	for n < len(nums) {
		if nums[i] != nums[n] {
			i++
			nums[i] = nums[n]
		}
		n++
	}
	return i + 1
}

// 这个划分感觉不会有任何效果，因为只是画成了k个组，而k个组都是O(n/k)的
// 累计起来之后还是相当于是O(n)，所以没什么优化效果
func removeDuplicates_version2(nums []int) int {
	return removeDuplicates_version2R(nums, 0, len(nums))
}

func removeDuplicates_version2R(nums []int, start, end int /* 不包含end */) int {
	if nums == nil || start == end {
		return 0
	}
	if (end - start) == 1 {
		return 1
	} else if (end - start) == 2 {
		if nums[start] == nums[start+1] {
			return 1
		} else {
			return 2
		}
	} else {
		mid := (start+end)/2 + 1
		leftN := removeDuplicates_version2R(nums, start, mid)
		rightN := removeDuplicates_version2R(nums, mid, end)
		// 左半部分和右半部分最多只可能有一个数字重复（且左边为最后一个，右边为第一个）
		if nums[start+leftN-1] == nums[mid] {
			rightN--
			mid++
		}
		n := 0
		for n < rightN {
			nums[start+leftN+n] = nums[mid+n]
			n++
		}
		return leftN + rightN
	}
}
