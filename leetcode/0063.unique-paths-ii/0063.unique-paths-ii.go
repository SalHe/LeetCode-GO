package leetcode

const (
	// nonObstacle = 0
	obstacle = 1
)

func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// 求格子形状
	m := len(obstacleGrid)    // >=1
	n := len(obstacleGrid[0]) // >=1

	// 分配dp
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	// 初始化dp
	if obstacleGrid[0][0] == obstacle {
		return 0
	}
	dp[0][0] = 1
	for i := 1; i < n; i++ {
		if obstacleGrid[0][i] == obstacle {
			break
		}
		dp[0][i] = dp[0][i-1]
	}
	for i := 1; i < m; i++ {
		if obstacleGrid[i][0] == obstacle {
			break
		}
		dp[i][0] = 1
	}

	// 开始求解
	for y := 1; y < m; y++ {
		for x := 1; x < n; x++ {
			if obstacleGrid[y][x] == obstacle {
				dp[y][x] = 0
			} else {
				dp[y][x] = dp[y-1][x] + dp[y][x-1]
			}
		}
	}

	return dp[m-1][n-1]
}
