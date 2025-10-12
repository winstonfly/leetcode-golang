package matrix

// NO.54 1行，i列，i行，1列
func spiralOrder(matrix [][]int) []int {
	var ans []int
	visited := make([][]bool, len(matrix))
	for i := range visited {
		visited[i] = make([]bool, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			visited[i][j] = false
		}
	}

	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右，下，左，上
	rows, cols := len(matrix), len(matrix[0])

	count := 0
	//从0, 0开始向右，遇到边界或者已经访问过的元素就转向
	for i, j := 0, 0; i <= rows && j <= cols; {
		if len(ans) == rows*cols {
			break
		}
		dr := directions[count%4]
		ni, nj := i+dr[0], j+dr[1]
		if i >= 0 && i < rows && j >= 0 && j < cols && !visited[i][j] {
			ans = append(ans, matrix[i][j])
			visited[i][j] = true
		}

		if ni < 0 || ni >= rows || nj < 0 || nj >= cols || visited[ni][nj] {
			count++
			continue
		}

		i, j = ni, nj
	}
	return ans
}
