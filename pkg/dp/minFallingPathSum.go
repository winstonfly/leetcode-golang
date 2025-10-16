package dp

// NO.931
func minFallingPathSum(matrix [][]int) int {
	var ans int
	// dp[i][j]表示到达第i行j列最小路径和，则有: dp[i][j] = min(dp[i-1][j-1], dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	dp[0][0] = matrix[0][0]
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[i][j] = matrix[i][j]
				continue
			}

			if j+1 < n {
				if j == 0 {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j+1]) + matrix[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i-1][j+1])) + matrix[i][j]
				}
			} else {
				if j == 0 {
					dp[i][j] = dp[i-1][j] + matrix[i][j]
				} else {
					dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + matrix[i][j]
				}
			}

		}
	}

	for i := 0; i < n; i++ {
		if i == 0 {
			ans = dp[m-1][i]
		} else {
			ans = min(ans, dp[m-1][i])
		}
	}

	return ans
}
