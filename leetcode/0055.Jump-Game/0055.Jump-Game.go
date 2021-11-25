package leetcode

func canJump(nums []int) bool {
	right := 0
	target := len(nums) - 1
	for i, num := range nums {
		if i > right {
			return false
		}
		nowRight := i + num
		if nowRight > right {
			right = nowRight
		}
		if right >= target {
			return true
		}
	}
	return false
}
