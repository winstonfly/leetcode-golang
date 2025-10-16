package dp

// NO.63
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// dp[i][i]表示达到i, j 的路径数, 则dp[i][j] = dp[i-1][j] + dp[i][j-1]
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		if obstacleGrid[i][0] == 1 {
			break
		}

		dp[i][0] = 1
	}

	for i := 0; i < n; i++ {
		if obstacleGrid[0][i] == 1 {
			break
		}
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 1 {
				dp[i][j] = 0
			} else {
				if i >= 1 {
					dp[i][j] += dp[i-1][j]
				}
				if j >= 1 {
					dp[i][j] += dp[i][j-1]
				}
			}

		}
	}

	return dp[m-1][n-1]
}
