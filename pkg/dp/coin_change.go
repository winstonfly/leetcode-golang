package dp

import "math"

// NO.322
func coinChange(coins []int, amount int) int {
	//i 代表金额为i时所需要的最小硬币数，最后一个硬币为c时，则dp[i] = dp[i-c] + 1， 最后一个硬币的可能性要从coins里取，所以有多个dp[i]，取最小的那个即可
	//定义dp的个数
	dp := make([]int, amount+1)
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		minn := math.MaxInt32
		for _, coin := range coins {
			if i >= coin {
				minn = min(minn, dp[i-coin])
			}
		}

		dp[i] = minn + 1
	}

	return dp[amount]
}
