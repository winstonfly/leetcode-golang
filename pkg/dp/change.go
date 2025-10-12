package dp

// NO.518
// 本题是求组合数，而零钱兑换1是求最少硬币数， 如果先遍历金额再遍历硬币，则会重复计算组合数
func change(amount int, coins []int) int {

	//dp[i]代表凑成金额i的组合数 dp[i] = dp[i] + dp[i-coin]
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] += dp[i-c]
		}
	}
	return dp[amount]
}
