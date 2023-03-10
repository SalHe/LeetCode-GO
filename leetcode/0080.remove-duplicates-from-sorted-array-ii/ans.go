package leetcode

func removeDuplicates(nums []int) int {
	// 有序，空间O(1)
	// 要求：元素最多出现两次
	length := len(nums)
	for i := len(nums) - 1; i >= 0; {
		num := nums[i]
		j := i - 1
		for ; j >= 0 && nums[j] == num; j-- {
		}
		exceeded := i - j - 2
		if exceeded > 0 {
			moved := length - i + 1
			copy(nums[j+1:j+1+moved], nums[i-1:i-1+moved])
			length -= exceeded
		}
		i = j

	}
	return length
}
