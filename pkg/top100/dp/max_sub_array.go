package dp

func maxSubArray(nums []int) int {
	//dp[i]代表以i结尾子数组的最大和， 则dp[i] = max(dp[i-1] + nums[i], nums[i]), 将所有i结尾的最大子数组求出保存，然后取出数据组最大的数即为最大。
	// dp[i] = max(dp[i-1] + nums[i], nums[i])
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}

	return ans
}
