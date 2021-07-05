package leetcode

func divide(dividend int, divisor int) int {
	// 不得使用乘法、除法、MOD

	// 先处理特殊情况：
	// 1. 被除数为0
	// 2. 除数与被除数绝对值相等
	// 3. 除数为1或-1
	// 4. 为2的倍数
	// 5. 为2的幂
	if dividend == 0 {
		return 0
	}
	if dividend == divisor {
		return 1
	}
	if dividend == -divisor {
		return -1
	}

	// 是2的倍数
	if divisor&1 == 0 {
		cur, power := divisor>>1, 1
		if cur&1 == 0 {
			power++
			cur >>= 1
		}

		if cur == 1 {
			// 是2的幂
			negative := (dividend > 0 && divisor < 0) || (dividend < 0 && divisor > 0)
			if dividend < 0 {
				dividend = -dividend
			}
			if divisor < 0 {
				divisor = -divisor
			}

			if negative {
				return -(dividend >> power)
			} else {
				return dividend >> power
			}
		} else if cur > 1 {
			// 仅仅只是2的倍数
			// 可以均除以2减少后面的计算量
			divisor >>= 1
			dividend >>= 1
		}
	}

	var factor int32
	if divisor == 1 {
		if dividend <= -2147483648 {
			factor = -2147483648
		} else if dividend >= 2147483647 {
			factor = 2147483647
		}
	} else if divisor == -1 {
		dividend = -dividend
		if dividend <= -2147483648 {
			factor = -2147483648
		} else if dividend >= 2147483647 {
			factor = 2147483647
		}
	} else {
		positive := true
		if (divisor > 0 && dividend < 0) || (divisor < 0 && dividend > 0) {
			positive = false
		}
		// 都转成负数来计算吧
		if divisor > 0 {
			divisor = -divisor
		}
		if dividend > 0 {
			dividend = -dividend
		}

		if divisor < dividend {
			return 0
		}

		factor = 1
		curr := divisor
		for curr+divisor >= dividend {
			if curr+curr >= dividend {
				curr += curr
				factor += factor
			} else {
				curr += divisor
				factor++
			}

			if curr == dividend {
				break
			}
		}

		if !positive {
			factor = -factor
		}
	}

	return int(factor)
}
