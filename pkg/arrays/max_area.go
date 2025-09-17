package arrays

// NO.11
func maxArea(height []int) int {

	left, right := 0, len(height)-1

	//公式: (right - left) * min(height[left], height[right])

	var ans int
	for left < right {
		ans = max(ans, (right-left)*min(height[left], height[right]))
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return ans
}
