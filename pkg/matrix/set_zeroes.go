package matrix

// NO.73
func setZeroes(matrix [][]int) {

	var zeros [][]int

	row, col := len(matrix), len(matrix[0])
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == 0 {
				zeros = append(zeros, []int{i, j})
			}
		}
	}

	for a := range zeros {
		k, v := zeros[a][0], zeros[a][1]
		for i := 0; i < col; i++ {
			matrix[k][i] = 0
		}

		for j := 0; j < row; j++ {
			matrix[j][v] = 0
		}
	}
}
