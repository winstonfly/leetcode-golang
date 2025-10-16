package dp

// NO.120
func minimumTotal(triangle [][]int) int {
	// dp[i][j]表示到达第i行j列的最小路径和 dp[i][j] = min(dp[i-1][j], d[[i-1][j-1]) + triangle[i][j]
	dp := make([][]int, len(triangle))
	m := len(triangle)
	for i := range dp {
		dp[i] = make([]int, len(triangle[i]))
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {

			maxQueue := len(triangle[i-1]) - 1
			if j == 0 {
				dp[i][j] = dp[i-1][j] + triangle[i][j]
				continue
			} else if j > maxQueue {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			} else {
				dp[i][j] = min(dp[i-1][j], dp[i-1][j-1]) + triangle[i][j]
			}
		}
	}

	var ans = 1 << 31
	for i := 0; i < len(dp[m-1]); i++ {
		ans = min(ans, dp[m-1][i])
	}
	return ans
}
