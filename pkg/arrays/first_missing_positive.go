package arrays

// NO.41  不符合题解， 题目要求用常量级的空间
func firstMissingPositive(nums []int) int {
	//第一遍，将小于0的标记为N=1
	n := len(nums)
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	//将小于n的标记为负数
	for i := 0; i < n; i++ {
		num := abs(nums[i])
		if num <= n {
			nums[i] = -nums[i]
		}
	}

	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i
		}
	}

	return n + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
