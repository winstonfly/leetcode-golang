package search

// NO.34
func searchRange(nums []int, target int) []int {
	if len(nums) == 1 {
		if nums[0] == target {
			return []int{0, 0}
		}
		return []int{-1, -1}
	}
	ans := []int{-1, -1}
	left, right := 0, len(nums)-1
	for left <= len(nums)-1 && right >= 0 {
		if nums[left] == target && ans[0] == -1 {
			ans[0] = left
		}

		if nums[right] == target && ans[1] == -1 {
			ans[1] = right
		}

		left++
		right--

	}

	if len(ans) == 0 {
		return []int{-1, -1}
	}

	return ans
}
