package dp

// NO.377
func combinationSum4(nums []int, target int) int {

	//dp[i]代和为i的组合数， 则dp[i] += dp[i-num] 完全背包

	dp := make([]int, target+1)
	dp[0] = 1
	for i := 0; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}

	return dp[target]
}
