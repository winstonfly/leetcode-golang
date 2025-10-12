package arrays

// NO.238
func productExceptSelf(nums []int) []int {
	left := make([]int, len(nums))
	right := make([]int, len(nums))

	left[0] = 1
	for i := 1; i < len(nums); i++ {
		left[i] = nums[i-1] * left[i-1]
	}

	right[len(nums)-1] = 1
	for i := len(nums) - 2; i >= 0; i-- {
		right[i] = nums[i+1] * right[i+1]
	}

	ans := make([]int, len(nums))

	for i := 0; i < len(nums); i++ {
		ans[i] = left[i] * right[i]
	}

	return ans
}
