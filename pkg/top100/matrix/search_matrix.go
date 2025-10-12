package matrix

// NO.240
func searchMatrix(matrix [][]int, target int) bool {
	for row := range matrix {
		if search(matrix[row], target) {
			return true
		}
	}

	return false
}

func search(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	mid := len(nums) / 2
	if nums[mid] == target {
		return true
	}

	if nums[mid] > target {
		return search(nums[:mid], target)
	} else {
		return search(nums[mid+1:], target)
	}

	return false
}
