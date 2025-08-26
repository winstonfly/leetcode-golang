package dp

import "math"

// NO.279
func numSquares(n int) int {
	// 当n=i时，平方数的取值在{1,j}之间, 如果dp[i]标记为值为i时平方数的最小个数, 平方数最后一个为c时，则dp[i] = dp[i-c] +1
	// 由于c是从平方数列表中取出，所以c需要枚举一下，并取出最小的最小平方数的值，即: dp[i] = min(dp[i-c]) + 1
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		minn := math.MaxInt32
		for j := 1; j*j <= i; j++ {
			minn = min(minn, dp[i-j*j])
		}

		dp[i] = minn + 1
	}
	return dp[n]
}
