package dp

// NO.198
func rob(nums []int) int {
	//最大值要嘛偷到最后一间，要嘛不偷 dp[i] = max(dp[i-1], dp[i-2] + nums[i])
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	dp := make([]int, len(nums))
	dp[0], dp[1] = nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(dp)-1]
}
