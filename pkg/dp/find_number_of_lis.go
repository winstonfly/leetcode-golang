package dp

// NO.673
func findNumberOfLIS(nums []int) int {
	var ans int

	//dp[i]表示以nums[i]结尾时最长递增子序列的长度
	dp := make([]int, len(nums))
	//count[i]表示以nums[i]结尾时最长递增子序列的个数
	count := make([]int, len(nums))
	maxLen := 0
	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		count[i] = 1

		// 1,3,5,4,7
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
					count[i] = count[j] // 重新计数
				} else if dp[j]+1 == dp[i] {
					count[i] += count[j]
				}
			}
		}

		if dp[i] > maxLen {
			maxLen = dp[i]
			ans = count[i]
		} else if dp[i] == maxLen {
			ans += count[i]
		}
	}

	return ans
}
