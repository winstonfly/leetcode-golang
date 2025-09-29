package search

// NO.34
func searchRange(nums []int, target int) []int {
	var ans []int
	leftIndex, rightIndex := -1, -1
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			//向左找
			for i := mid; i >= 0; i-- {
				if nums[i] == target {
					leftIndex = i
				} else {
					break
				}
			}

			//向右找
			for i := mid; i < len(nums); i++ {
				if nums[i] == target {
					rightIndex = i
				} else {
					break
				}
			}
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	ans = append(ans, leftIndex, rightIndex)
	return ans
}
