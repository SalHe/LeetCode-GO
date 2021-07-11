package leetcode

func searchRange(nums []int, target int) []int {
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
