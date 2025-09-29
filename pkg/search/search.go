package search

// NO.33
func search(nums []int, target int) int {
	ans := -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			ans = mid
			return ans
		}

		if nums[mid] >= nums[left] {
			//左边有序
			if target >= nums[left] && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			//右边有序
			if target > nums[mid] && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}

	return ans
}
