package backtrack

// NO.79
func exist1(board [][]byte, word string) bool {
	var ans bool
	m, n := len(board), len(board[0])

	dp := make([][]bool, m)
	for i := range dp {
		dp[i] = make([]bool, n)
	}
	dr, dc := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	var backtrack func(i, j int, curr string)
	backtrack = func(i, j int, curr string) {
		if ans {
			return
		}

		if curr == word {
			ans = true
			return
		}

		for k := 0; k < 4; k++ {
			x, y := i+dr[k], j+dc[k]
			if x >= 0 && x < m && y >= 0 && y < n && !dp[x][y] && board[x][y] == word[len(curr)] {
				dp[x][y] = true
				curr += string(board[x][y])
				backtrack(x, y, curr)
				dp[x][y] = false
				curr = curr[:len(curr)-1]
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				dp[i][j] = true
				backtrack(i, j, string(board[i][j]))
				dp[i][j] = false
			}
			if ans {
				return ans
			}
		}
	}
	return ans
}

func exist(board [][]byte, word string) bool {
	var ans bool

	m, n := len(board), len(board[0])
	visited := make([][]bool, m)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	dr, dc := []int{-1, 1, 0, 0}, []int{0, 0, -1, 1}
	var backtrack func(i, j int, word string)
	backtrack = func(i, j int, word string) {
		if board[i][j] != word[0] || visited[i][j] {
			return
		}

		if len(word) == 1 && board[i][j] == word[0] {
			ans = true
			return
		}

		for d := 0; d < 4; d++ {
			x, y := i+dr[d], j+dc[d]
			if x >= 0 && x < m && y >= 0 && y < n && !visited[x][y] {
				visited[i][j] = true
				backtrack(x, y, word[1:])
				visited[i][j] = false
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == word[0] {
				backtrack(i, j, word)
			}
			if ans {
				return ans
			}
		}
	}

	return ans
}
