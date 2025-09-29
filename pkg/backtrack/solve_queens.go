package backtrack

// NO.51 N皇后问题 经典回溯法解决
func solveNQueens(n int) [][]string {
	var ans [][]string

	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
	}

	col := make([]bool, n)
	diag1 := make([]bool, 2*n-1) //标记左上到右下(i+j)方向的对角线
	diag2 := make([]bool, 2*n-1) //标记右上到左下(i-j+n-1)方向的对角线

	var backtrack func(i int)
	backtrack = func(i int) {
		if i == n {
			ans = append(ans, generateBoard(dp))
			return
		}

		for j := 0; j < n; j++ {
			if col[j] {
				continue
			}

			if diag1[i+j] {
				continue
			}

			if diag2[i-j+n-1] {
				continue
			}

			dp[i][j] = true
			col[j] = true
			diag1[i+j] = true
			diag2[i-j+n-1] = true
			backtrack(i + 1)
			dp[i][j] = false
			col[j] = false
			diag1[i+j] = false
			diag2[i-j+n-1] = false
		}
	}

	backtrack(0)

	return ans
}

func generateBoard(dp [][]bool) []string {
	var board []string
	for i := range dp {
		row := make([]byte, len(dp))
		for j := range dp[i] {
			if dp[i][j] {
				row[j] = 'Q'
			} else {
				row[j] = '.'
			}
		}

		board = append(board, string(row))
	}

	return board
}
