package dp

// NO.322
func coinChange(coins []int, amount int) int {
	// 01背包问题 dp[i]代表凑成金额i所需要的最少硬币数，dp[i] = min(dp[i], dp[i-coin] + 1)

	dp := make([]int, amount+1)
	dp[0] = 0

	for i := 1; i < amount+1; i++ {
		dp[i] = amount + 1
		for _, c := range coins {
			if i >= c {
				dp[i] = min(dp[i], dp[i-c]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}

	return dp[amount]
}
