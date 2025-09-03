package dp

// NO.70
func climbStairs(n int) int {

	// n = 2时有两种方法到楼顶， dp[i] 设到第n层时的方法量，则dp[i] =dp[i-1]  + dp[i-2]
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}

	return dp[n]
}
