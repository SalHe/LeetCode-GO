package leetcode

import "math"

func myAtoi(s string) int {
	number := 0
	hasStarted := false
	hasSigned := false
	isNegative := false

	for _, c := range s {
		if c >= '0' && c <= '9' {
			num := int(c - '0')

			// 溢出检测
			if !isNegative && math.MaxInt32-number*10 < num {
				number = math.MaxInt32
			} else if isNegative && math.MaxInt32+1-number*10 < num {
				number = math.MaxInt32 + 1
			} else {
				number = number*10 + num
			}

			hasStarted = true
		} else if c == '-' || c == '+' {
			if hasStarted {
				break
			}
			if hasSigned {
				return 0
			}
			isNegative = c == '-'
			hasSigned = true
			hasStarted = true
		} else if c == ' ' {
			if hasStarted {
				break
			}
			continue
		} else {
			break
		}
	}

	if isNegative {
		return -number
	}
	return number
}
