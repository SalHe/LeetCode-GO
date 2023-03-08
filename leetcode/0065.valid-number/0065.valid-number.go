package leetcode

// 这个题还可以用状态机做
func isNumber(s string) bool {
	// 1 <= s.length <= 20
	i := 0
	if s[i] == '+' || s[i] == '-' {
		i++
	}

	integral := false
	for i < len(s) && s[i] >= '0' && s[i] <= '9' {
		integral = true
		i++
	}
	if i >= len(s) {
		return integral
	}

	fractional := false
	if s[i] == '.' {
		i++
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			fractional = true
			i++
		}
	}

	if !integral && !fractional {
		return false
	}
	if i >= len(s) {
		return true
	}

	if (s[i] == 'e' || s[i] == 'E') && i < len(s)-1 {
		i++
		if (s[i] == '+' || s[i] == '-') && i < len(s)-1 {
			i++
		}
		for i < len(s) && s[i] >= '0' && s[i] <= '9' {
			i++
		}
	}

	return i == len(s)
}
