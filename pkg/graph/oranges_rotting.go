package graph

// NO.994 广度优先搜索
func orangesRotting(grid [][]int) int {
	R, C := len(grid), len(grid[0])
	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}
	queue := []int{}
	depth := map[int]int{}

	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, i*C+j)
				depth[i*C+j] = 0
			}
		}
	}

	ans := 0
	for len(queue) > 0 {
		code := queue[0]
		queue = queue[1:]
		r, c := code/C, code%C

		for k := 0; k < 4; k++ {
			nr, nc := r+dr[k], c+dc[k]
			if nc < 0 || nc >= C || nr < 0 || nr >= R || grid[nr][nc] != 1 {
				break
			}

			grid[nr][nc] = 2
			ncode := nr*C + nc
			queue = append(queue, ncode)
			depth[ncode] = depth[code] + 1
			ans = depth[ncode]
		}
	}

	for _, v := range grid {
		for _, vv := range v {
			if vv == 1 {
				return -1
			}
		}
	}

	return ans
}
