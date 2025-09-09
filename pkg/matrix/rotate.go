package matrix

// NO.48
func rotate(matrix [][]int) [][]int {
	n := len(matrix)
	tmp := make([][]int, n)
	for i := range tmp {
		tmp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			tmp[i][j] = matrix[n-j-1][i]
		}
	}

	copy(matrix, tmp)
	return matrix
}
