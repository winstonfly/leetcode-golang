package greed

// NO.55
func canJump(nums []int) bool {
	lastPos := len(nums) - 1
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i]+i >= lastPos {
			lastPos = i
		}
	}

	if lastPos == 0 {
		return true
	}

	return false
}
