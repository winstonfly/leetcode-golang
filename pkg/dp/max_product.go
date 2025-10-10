package dp

// NO.152
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	fMax := make([]int, len(nums))
	fMin := make([]int, len(nums))
	var ans int

	fMax[0], fMin[0], ans = nums[0], nums[0], nums[0]

	for i := 1; i < len(nums); i++ {
		fMax[i] = max(max(fMax[i], fMax[i-1]*nums[i]), max(nums[i], fMin[i-1]*nums[i]))
		fMin[i] = min(min(fMin[i], fMin[i-1]*nums[i]), min(nums[i], fMax[i-1]*nums[i]))
	}

	for _, v := range fMax {
		ans = max(ans, v)
	}

	return ans
}
