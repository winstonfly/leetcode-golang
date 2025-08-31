package matrix

// NO.73
func setZeroes(matrix [][]int) {

	dp := make([][]bool, len(matrix))
	for i := range dp {
		dp[i] = make([]bool, len(matrix[0]))
	}
	row, col := len(matrix), len(matrix[0])

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if dp[i][j] {
				for k := 0; k < row; k++ {
					matrix[k][j] = 0
				}
				for k := 0; k < col; k++ {
					matrix[i][k] = 0
				}
			}
		}
	}
}
