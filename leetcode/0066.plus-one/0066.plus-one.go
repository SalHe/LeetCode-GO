package leetcode

func plusOne(digits []int) []int {
	// 1 <= digits.length <= 100
	count := len(digits)
	if digits[count-1] != 9 {
		digits[count-1]++
		return digits
	}

	i := count - 1
	for i >= 0 && digits[i] == 9 {
		digits[i] = 0
		i--
	}

	if i <= -1 {
		return append([]int{1}, digits...)
	} else {
		digits[i]++
		return digits
	}
}
