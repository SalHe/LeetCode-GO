package leetcode

func sortColors(nums []int) {
	ps := [3]int{0, 0, 0}
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		if n <= 1 {
			copy(nums[ps[n]+1:i+1], nums[ps[n]:i])
			nums[ps[n]] = n

		}
		for j := n; j < len(ps); j++ {
			ps[j]++
		}
	}
}
