package leetcode

func mySqrt(x int) int {
	if x == 0 {
		return 0
	}

	min, max := 1, x
	// 针对4的倍数优化
	for max&0b11 == 0 {
		min *= 2
		max >>= 1
	}
	if max < min {
		max, min = min, max
	}

	for min < max {
		mid := (max-min+1)/2 + min
		if x/mid >= mid {
			min = mid
		} else {
			max = mid - 1
		}
	}

	return min
}
