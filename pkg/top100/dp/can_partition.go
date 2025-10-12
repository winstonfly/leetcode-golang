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
	//dp[i]表示容量为i的背包是否能被填满
	dp := make([]bool, total+1)
	dp[0] = true
	// 1， 2， 5， total = 5
	for i := 0; i <= len(nums)-1; i++ {
		for j := 0; j < total+1; j++ {
			if i >= nums[i] {
				dp[j] = dp[j] || dp[j-nums[i]]
			}
		}
	}

	return dp[total]
}
