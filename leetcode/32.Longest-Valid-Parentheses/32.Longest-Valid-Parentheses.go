package leetcode

func longestValidParentheses(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}

	a := make([][]bool, n)
	max, maxTails := make([]int, n), make([]int, n)
	for i := range a {
		a[i] = make([]bool, (n-i)/2+1)
		a[i][0] = true
	}

	maxPairs := n / 2
	result := 0
	for pl := 1; pl <= maxPairs; pl += 1 {
		found := false
		for i := 0; i < n-1; i++ {
			if pl >= len(a[i]) {
				break
			}
			a[i][pl] = (s[i] == '(' && s[i+pl*2-1] == ')' && a[i+1][pl-1]) || // (....)
				(pl >= 2 && a[i][1] && a[i+2][pl-2] && a[i+pl*2-2][1]) // ()....() , .... 可以是空串
			if !(a[i][pl]) && pl >= 3 && s[i] == '(' && s[i+pl*2-1] == ')' {

				max1 := max[i]
				if max1 > 0 && maxTails[i+pl*2-1] > 0 {
					endpos := i + pl*2 - 1
					start2 := endpos - maxTails[i+pl*2-1]*2 + 1
					s2 := start2
					pl2 := (endpos - s2 + 1) / 2
					// 前面个条件等我再详细证明下，粗略地看是可以的
					a[i][pl] = pl-max1-pl2 < 0 || a[i][max1] && a[i+max1*2][pl-max1-pl2] && a[s2][pl2]
				}
				// for s2 := start2; s2 >= i+max1*2; s2 -= 2 {
				// 	pl2 := (endpos - s2 + 1) / 2
				// 	if !(a[i][pl]) &&
				// 		a[i][max1] && a[i+max1*2][pl-max1-pl2] && a[s2][pl2] {
				// 		a[i][pl] = true
				// 		break
				// 	}
				// }
			}
			if a[i][pl] {
				if pl > result {
					result = pl
				}
				found = true
				max[i] = pl
				maxTails[i+pl*2-1] = pl
			}
		}

		// 如果有效非空括号串存在
		// 那么至少有一对有效括号
		// 但是可能存在有n对有效括号，却无n-1对有效括号的情况
		// 比如 (())(())存在4对有效括号的串，但是找不到3对有效括号的串
		if pl == 1 && !found {
			return 0
		}
	}

	return result * 2
}

func longestValidParentheses_version1(s string) int {
	n := len(s)
	if n <= 0 {
		return 0
	}
	a := make([][]bool, n)
	max := make([]int, n)
	for i := range a {
		a[i] = make([]bool, (n-i)/2+1)
		a[i][0] = true
	}

	maxPairs := n / 2
	result := 0
	for pl := 1; pl <= maxPairs; pl += 1 {
		for i := 0; i < n-1; i++ {
			if pl >= len(a[i]) {
				break
			}
			a[i][pl] = (s[i] == '(' && s[i+pl*2-1] == ')' && a[i+1][pl-1]) || // (....)
				(pl >= 2 && a[i][1] && a[i+2][pl-2] && a[i+pl*2-2][1]) // ()....() , .... 可以是空串
			if !(a[i][pl]) && pl >= 3 {

				max1 := max[i]

				endpos := i + pl*2 - 1 // included
				start2 := endpos - 1
				for s2 := start2; s2 >= i+max1*2; s2 -= 2 {
					pl2 := (endpos - s2 + 1) / 2
					if !(a[i][pl]) &&
						a[i][max1] && a[i+max1*2][pl-max1-pl2] && a[s2][pl2] {
						a[i][pl] = true
						break
					}
				}
			}
			if a[i][pl] {
				if pl > result {
					result = pl
				}
				max[i] = pl
			}
		}
	}

	return result * 2
}
