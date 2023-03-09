package leetcode

func minDistance(word1 string, word2 string) int {
	// 0 <= word1.length, word2.length <= 500
	m := len(word1)
	n := len(word2)
	// if m < n {
	// 	word1, word2 = word2, word1
	// 	m, n = n, m
	// }
	if m == 0 {
		return n
	}
	if n == 0 {
		return m
	}

	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	if word1[0] != word2[0] {
		dp[0][0] = 1
	}
	for i := 1; i < m; i++ {
		if word1[i] == word2[0] {
			dp[i][0] = i
		} else {
			dp[i][0] = min(i, dp[i-1][0]) + 1
		}
	}
	for j := 1; j < n; j++ {
		if word1[0] == word2[j] {
			dp[0][j] = j
		} else {
			dp[0][j] = min(j, dp[0][j-1]) + 1
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ch1, ch2 := word1[i], word2[j]
			if ch1 == ch2 {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(min(dp[i][j-1], dp[i-1][j]), dp[i-1][j-1]) + 1
			}
		}
	}

	return dp[m-1][n-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
