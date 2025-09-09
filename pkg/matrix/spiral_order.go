package matrix

// NO.54 1行，i列，i行，1列
func spiralOrder(matrix [][]int) []int {

	rows, cols := len(matrix), len(matrix[0])

	dp := make([][]bool, rows)
	for i := range dp {
		dp[i] = make([]bool, cols)
	}

	var ans []int
	var order func(i, j, direction int)
	//direction: 0 向右 1 向下  2 向左 3 向上
	order = func(row, col, direction int) {
		if len(ans) == rows*cols {
			return
		}
		switch direction {
		case 0:
			for col < cols && !dp[row][col] {
				ans = append(ans, matrix[row][col])
				dp[row][col] = true
				col++
			}
			order(row+1, col-1, direction+1)
		case 1:
			for row < rows && !dp[row][col] {
				ans = append(ans, matrix[row][col])
				dp[row][col] = true
				row++
			}

			order(row-1, col-1, direction+1)

		case 2:
			for col >= 0 && !dp[row][col] {
				ans = append(ans, matrix[row][col])
				dp[row][col] = true
				col--
			}

			order(row-1, col+1, direction+1)
		case 3:
			for row >= 0 && !dp[row][col] {
				ans = append(ans, matrix[row][col])
				dp[row][col] = true
				row--
			}
			order(row+1, col+1, 0)
		}
	}

	order(0, 0, 0)

	return ans
}
