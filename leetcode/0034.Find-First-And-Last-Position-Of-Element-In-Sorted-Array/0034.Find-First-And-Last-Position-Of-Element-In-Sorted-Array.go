package leetcode

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	return searchRange_binary(nums, target, 0, len(nums)-1)
}

func searchRange_binary(nums []int, target int, start int, end int) []int {
	if start > end {
		return []int{-1, -1}
	}
	mid := start + (end-start)/2
	if nums[mid] == target {
		// 分别向左右找边界
		s, e := mid, mid

		// 寻找左边界（最靠右且小于[mid]的数）
		left, right := start, mid
		center := left + (right-left)/2
		for left <= right {
			if nums[center] == target {
				s = center
				if center == left || nums[center-1] < target {
					break
				}
				right = center - 1
			} else if center < right && nums[center+1] == target {
				s = center + 1
				break
			} else {
				left = center + 1
			}
			center = left + (right-left)/2
		}
		// 寻找右边界（最靠左且大于[mid]的数）
		left, right = mid, end
		center = left + (right-left)/2
		for left <= right {
			if nums[center] == target {
				e = center
				if center == right || nums[center+1] > target {
					break
				}
				left = center + 1
			} else if center > left && nums[center-1] == target {
				e = center - 1
				break
			} else {
				right = center - 1
			}
			center = left + (right-left)/2
		}

		return []int{s, e}
	} else if nums[mid] > target {
		// 说明target位于左边部分的数组
		return searchRange_binary(nums, target, start, mid-1)
	} else {
		// 说明target位于右边部分的数组
		return searchRange_binary(nums, target, mid+1, end)
	}
}

func searchRange_linear(nums []int, target int) []int {
	start, end := -1, -1
	for i := 0; i < len(nums); i++ {
		if nums[i] > target {
			break
		} else if nums[i] == target {
			if start == -1 {
				start = i
			}
			end = i
		}
	}
	return []int{start, end}
}
