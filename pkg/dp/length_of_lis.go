package dp

// NO.300
func lengthOfLIS(nums []int) int {

	// f[i] 代表以nums[i]结尾的最长递增子序列的长度
	f := make([]int, len(nums))
	f[0] = 1

	for i := 1; i < len(nums); i++ {
		f[i] = 1
		for j := 0; j <= i; j++ {
			if nums[i] > nums[j] {
				f[i] = max(f[i], f[j]+1)
			}
		}

	}

	var ans int
	for i := 0; i < len(f); i++ {
		ans = max(ans, f[i])
	}

	return ans
}
