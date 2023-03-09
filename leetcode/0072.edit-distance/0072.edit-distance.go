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

	// DP分配
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化DP数组（对首行、首列单独处理）
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

	// 开始DP过程
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			ch1, ch2 := word1[i], word2[j]
			if ch1 == ch2 {
				// "....a"
				// "____a"
				// 则步数为"...."与"____"的编辑距离

				dp[i][j] = dp[i-1][j-1]
			} else {
				// "....a"
				// "____b"
				// 1. 将'a'替换成'b'或'b'替换成'a'
				// 2. "____"变换到"....a"然后加上'b'
				// 3. 类似2

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
