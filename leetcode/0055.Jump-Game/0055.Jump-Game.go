package leetcode

func canJump(nums []int) bool {
	n := len(nums) // >=1
	dp := make([]bool, n)
	target := n - 1

	for i := n - 1; i >= 0; i-- {
		if i+nums[i] >= target {
			dp[i] = true
		} else {
			for dis := nums[i]; dis >= 1; dis-- {
				if dp[i+dis] {
					dp[i] = true
					break
				}
			}
		}
	}

	return dp[0]
}
