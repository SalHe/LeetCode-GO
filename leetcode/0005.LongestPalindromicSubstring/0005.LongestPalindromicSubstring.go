package leetcode

func longestPalindrome(s string) string {
	candidate := make([][]bool, len(s))
	for i := range candidate {
		candidate[i] = make([]bool, len(s))
	}

	longest := ""
	for subLen := 1; subLen <= len(s); subLen++ {
		for i := 0; i < len(s); i++ {
			j := i + subLen - 1
			if j >= len(s) {
				continue
			}
			if subLen <= 2 {
				candidate[i][j] = s[i] == s[j]
			} else {
				candidate[i][j] = s[i] == s[j] && i < len(s) && j > 0 && candidate[i+1][j-1]
			}

			if candidate[i][j] && len(longest) < subLen {
				longest = s[i : j+1]
			}
		}
	}
	return longest
}

func longestPalindrome_Manacher(s string) string {
	dr := make([]int, 2*len(s)+1)

	maxRight := 0
	center := 0
	maxLen := 1
	begin := 0
	for i := 0; i < 2*len(s)+1; i++ {
		if i < maxRight {
			mirror := 2*center - i
			dr[i] = min(dr[mirror], maxRight-i)
		}

		left, right := i-(dr[i]+1), i+(dr[i]+1)
		for left >= 0 && right < 2*len(s)+1 && ((left%2 == 0 && right%2 == 0) || s[(left-1)/2] == s[(right-1)/2]) {
			left--
			right++
			dr[i]++
		}

		if maxRight < i+dr[i] {
			maxRight = i + dr[i]
			center = i
		}

		if maxLen < dr[i] {
			maxLen = dr[i]
			begin = (i - dr[i]) / 2
		}
	}
	return s[begin : begin+maxLen]
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
