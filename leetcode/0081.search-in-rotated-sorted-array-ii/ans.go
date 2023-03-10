package leetcode

func search(nums []int, target int) bool {
	count := len(nums)
	if count == 2 {
		return target == nums[0] || target == nums[1]
	} else if nums[0] < nums[count-1] {
		return exists(nums, target)
	} else {
		g := count - 1 // 这部分应该是可以优化的
		for i := 0; i < count-1; i++ {
			if nums[i] > nums[i+1] {
				g = i + 1
				break
			}
		}
		return exists(nums[:g], target) || exists(nums[g:], target)
	}
}

func exists(nums []int, target int) bool {
	for left, right := 0, len(nums)-1; left <= right; {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return true
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
