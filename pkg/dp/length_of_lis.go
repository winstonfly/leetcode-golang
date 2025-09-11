package dp

// NO.300
func lengthOfLIS(nums []int) int {
	//动态规划解法，之前的想法是从前往后来获取dp[i]的值，保存前边的最后一个值，然后后边的大于该值就+1， 比较麻烦，其实应该换一种思路，直接
	//求所有以nums[i]结尾的最长的子数组序列长度，这样max(dp[i])就是最终的答案了
	dp := make([]int, len(nums))
	dp[0] = 1

	for i := 0; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
	}

	var ans int
	for _, v := range dp {
		ans = max(ans, v)
	}

	return ans
}
