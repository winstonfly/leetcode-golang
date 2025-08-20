package search

// NO.79
func exist(board [][]byte, word string) bool {
	var ans int
	if len(word) == 0 {
		return false
	}

	dp := make([][]bool, len(board))
	for i := range dp {
		dp[i] = make([]bool, len(board[0]))
	}

	var dfs func(board [][]byte, word string, i, j int)
	dfs = func(board [][]byte, word string, i, j int) {
		m := len(board)
		n := len(board[0])
		if i >= m || i < 0 || j >= n || j < 0 {
			return
		}
		if dp[i][j] {
			//已经使用过
			return
		}
		if len(word) == 0 {
			ans = ans + 1
			return
		}

		if board[i][j] != word[0] {
			return
		} else {
			//移动查询下一位
			dp[i][j] = true
			dfs(board, word[1:], i+1, j)
			dfs(board, word[1:], i-1, j)
			dfs(board, word[1:], i, j+1)
			dfs(board, word[1:], i, j-1)
			dp[i][j] = false
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			dfs(board, word, i, j)
		}
	}

	if ans > 0 {
		return true
	}
	return false
}
