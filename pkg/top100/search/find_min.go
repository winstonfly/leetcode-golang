package search

// 153. 寻找旋转排序数组中的最小值
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	var ans int
	for left < right {
		mid := left + (right-left)/2
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	ans = nums[left]

	return ans
}
