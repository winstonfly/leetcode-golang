package matrix

// NO.240
func searchMatrix(matrix [][]int, target int) bool {
	var check func(nums []int, target int) bool
	check = func(nums []int, target int) bool {
		if len(nums) == 1 && nums[0] == target {
			return true
		}
		mid := len(nums) / 2
		if mid == 0 {
			return false
		}

		if nums[mid] == target {
			return true
		}

		if target > nums[mid] {
			return check(nums[mid+1:], target)
		} else {
			return check(nums[0:mid], target)
		}
	}

	var ans bool
	for i := range matrix {
		if target >= matrix[i][0] && target <= matrix[i][len(matrix[i])-1] {
			if check(matrix[i], target) {
				ans = true
				break
			}
		}
	}

	return ans
}
