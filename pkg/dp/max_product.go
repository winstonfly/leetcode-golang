package dp

func maxProduct(nums []int) int {
	//dp[i]代表以第i个元素结尾时子数组乘积的最大值，求出所有结尾的值，那么这个dp数组里最大的那个即是该数组子数组的最大乘积
	//正常思维， dp[i] = max(dp[i-1] * nums[i], nums[i])， 但碰到有负数的情况要考虑进来， 也即是将负负得正的情况考虑进来
	maxF := make([]int, len(nums))
	minF := make([]int, len(nums))
	maxF[0] = nums[0]
	minF[0] = nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		maxF[i] = max(maxF[i-1]*nums[i], nums[i])
		minF[i] = min(maxF[i-1]*nums[i], min(minF[i-1]*nums[i], nums[i]))

		maxF[i] = max(maxF[i], minF[i-1]*nums[i])
		ans = max(ans, maxF[i])
	}

	return ans
}
