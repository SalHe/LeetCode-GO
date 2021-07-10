package leetcode

func search(nums []int, target int) int {
	// 原址实现
	n := len(nums)
	if n <= 0 {
		return -1
	}

	// 只有一个元素时特殊处理
	if n == 1 {
		if nums[0] == target {
			return 0
		} else {
			return -1
		}
	}

	// 由原数组的下标变换到旋转后数组下标所需要的增量（需要取模）
	k := 0
	for ; k < n; k++ {
		if k < n-1 && nums[k] > nums[k+1] {
			k++
			break
		}
	}

	// 二分查找
	left := 0
	right := n - 1
	mid := left + (right-left)/2
	for left <= right {
		if nums[(mid+k)%n] == target {
			return (mid + k) % n
		} else if nums[(mid+k)%n] > target {
			right = mid - 1
		} else { // nums[(mid+k)%n] < target
			left = mid + 1
		}
		mid = left + (right-left)/2
	}

	return -1
}
