package search

func searchMatrix(matrix [][]int, target int) bool {

	m := len(matrix)
	n := len(matrix[0])

	for i := 0; i < m; i++ {
		if matrix[i][n-1] < target {
			continue
		}

		left, right := 0, n-1
		for left <= right && right < n {
			if matrix[i][left] == target || matrix[i][right] == target {
				return true
			}

			if matrix[i][left] < target {
				left++
			}

			if matrix[i][right] > target {
				right--
			}
		}
		break
	}

	return false

}
