package leetcode

type dpNode struct {
	start int
	end   int
	valid bool
}

func (d *dpNode) distance() int {
	return d.end - d.start + 1
}

func (d *dpNode) extract(s string) string {
	if d.valid {
		return s[d.start : d.end+1]
	}
	return ""
}

func (d *dpNode) update(valid bool, start, end int) {
	if !d.valid || end-start+1 < d.distance() {
		d.start = start
		d.end = end
	}
	d.valid = valid
}

func (d *dpNode) updateFrom(node *dpNode) {
	d.update(node.valid, node.start, node.end)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minWindow_fail(s string, t string) string {

	// 根据题意 s t 均不是空串
	len1 := len(s)
	len2 := len(t)

	// dp分配
	dp := make([][]dpNode, len1)
	for i := 0; i < len1; i++ {
		dp[i] = make([]dpNode, len2)
	}

	// dp初始化
	dp[0][0].valid = s[0] == t[0]
	for y := 1; y < len1; y++ {
		if s[y] == t[0] {
			dp[y][0].update(true, y, y)
		} else {
			dp[y][0].updateFrom(&dp[y-1][0])
		}
	}

	// dp(将dp数组改成行列交换后按内存访问局部性应该可以提高性能)
	for x := 1; x < len2; x++ {
		for y := x; y < len1; y++ {
			if s[y] == t[x] {
				dp[y][x].update(dp[y-1][x-1].valid, min(dp[y-1][x-1].start, y), max(dp[y-1][x-1].end, y))
				if (y < len1-1 && dp[y][x].distance() > x+1) || !dp[y][x].valid {
					dp[y][x].update(dp[y+1][x-1].valid, min(dp[y+1][x-1].start, y), max(dp[y+1][x-1].end, y))
				}
			} else {
				// dp[y][x].updateFrom(&dp[y-1][x])
				for k := y - 1; k >= 0; k-- {
					if s[k] == t[x] {
						dp[y][x].update(dp[y][x-1].valid, min(dp[y][x-1].start, k), max(dp[y][x-1].end, k))
					}
				}
			}
		}
	}

	return dp[len1-1][len2-1].extract(s)
}
