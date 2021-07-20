package leetcode

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	dp := make([][]bool, ls+1)
	for i := 0; i < ls+1; i++ {
		dp[i] = make([]bool, lp+1)
	}
	// dp[x][0] 必定为false (x>=1)
	// 因为p的前0个字符（空串）与s的前x个字符组成的非空子串不匹配
	// dp[0][0] 应为true
	// 因为p的前0个字符（空串）与s的前0个字符组成的空子串匹配
	dp[0][0] = true

	// dp[0][x] 单独处理，简化下面的循环
	for j := 1; j < lp+1; j++ {
		dp[0][j] = dp[0][j-1] && p[j-1] == '*'
	}

	// 动态规划求dp[y][x] (x,y >= 1)
	for i := 1; i < ls+1; i++ {
		for j := 1; j < lp+1; j++ {
			ch1, ch2 := s[i-1], p[j-1]
			if ch2 == '?' || ch1 == ch2 {
				dp[i][j] = dp[i-1][j-1]
			} else if ch2 == '*' {
				dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
			}
		}
	}
	return dp[ls][lp]
}

// isMatch_dfs 有些用例无法在指定时间内完成
func isMatch_dfs(s string, p string) bool {
	var dfs func(s1, s2 int) bool
	dfs = func(s1, s2 int) bool {
		if s1 >= len(s) && s2 < len(p) {
			if s2 == len(p)-1 && p[s2] == '*' {
				return true
			} else if p[s2] == '*' {
				return dfs(s1+1, s2+1)
			}
			return false
		} else if s1 < len(s) && s2 >= len(p) {
			return false
		} else if s1 >= len(s) && s2 >= len(p) {
			return true
		}

		ch1, ch2 := s[s1], p[s2]
		if ch2 == '?' {
			return dfs(s1+1, s2+1)
		} else if ch2 == '*' && (s2 == 0 || p[s2-1] != '*') {
			if dfs(s1, s2+1) { // *不代表任何字符
				return true
			} else if dfs(s1+1, s2) { // *代表2个和多个字符
				return true
			} else if dfs(s1+1, s2+1) { // *代表1个字符
				return true
			} else {
				return false
			}
		} else if ch2 == '*' && s2 > 0 && p[s2-1] == '*' {
			return dfs(s1, s2+1)
		} else if ch2 == ch1 {
			return dfs(s1+1, s2+1)
		} else {
			return false
		}
	}
	return dfs(0, 0)
}
