package dp

// NO.416
func canPartition1(nums []int) bool {
	//dp[i]代表i个数组是否能拆成两个子数组，并且和相等
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	total := sum / 2
	//改为从nums里挑选出来的数据和为total
	var ans bool
	var backtrack func(start int, tmpSum int)
	backtrack = func(start int, tmpSum int) {

		if tmpSum == total {
			ans = true
			return
		}

		if tmpSum > total {
			return
		}

		for i := start; i < len(nums); i++ {
			tmpSum += nums[i]
			backtrack(i+1, tmpSum)
			tmpSum -= nums[i]
		}
	}

	backtrack(0, 0)

	return ans
}

func canPartition(nums []int) bool {
	//动态规划解法，将题目转为01背包问题
	sum := 0
	for _, v := range nums {
		sum += v
	}

	if sum%2 != 0 {
		return false
	}

	total := sum / 2
	dp := make([]bool, total+1)
	dp[0] = true

	//倒序枚举，保证不重复使用元素， 如果正序枚举会重复使用元素，成了完全背包问题了
	for i := 0; i < len(nums); i++ {
		for j := total; j >= nums[i]; j-- {
			dp[j] = dp[j] || dp[j-nums[i]]
		}
	}
	return dp[total]
}
