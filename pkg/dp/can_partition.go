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
	//01背包问题，动态规划
	sum, maxx := 0, 0

	for _, v := range nums {
		sum += v
		if v > maxx {
			maxx = v
		}
	}

	if sum%2 != 0 {
		return false
	}

	total := sum / 2
	if maxx > total {
		return false
	}

	//定义dp[i][j]表示前i个数据里是否能凑出和为j
	dp := make([][]bool, len(nums))
	for i, _ := range dp {
		dp[i] = make([]bool, total+1)
		dp[i][0] = true
	}

	dp[0][nums[0]] = true
	for i := 1; i < len(nums); i++ {
		for j := 1; j <= total; j++ {
			if j-nums[i] > 0 {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-nums[i]]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[len(nums)-1][total]
}
