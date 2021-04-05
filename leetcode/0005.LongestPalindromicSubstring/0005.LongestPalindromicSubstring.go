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
