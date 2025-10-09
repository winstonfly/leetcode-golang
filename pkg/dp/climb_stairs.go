package dp

// NO.70
func climbStairs(n int) int {
	//爬到第i阶可以从第i-1阶或者i-2阶爬上来 dp[i] = dp[i-1 + dp[i-2]
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
