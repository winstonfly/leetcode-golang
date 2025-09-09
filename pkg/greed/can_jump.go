package greed

// NO.55
func canJump(nums []int) bool {

	lastPos := len(nums) - 1
	for i := lastPos; i >= 0; i-- {
		if i+nums[i] >= lastPos {
			lastPos = i
		}
	}

	return lastPos == 0
}
