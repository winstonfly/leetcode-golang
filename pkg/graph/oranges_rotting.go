package graph

// NO.994 广度优先搜索
func orangesRotting(grid [][]int) int {
	var ans int
	rows := len(grid)
	cols := len(grid[0])

	//用于存储节点所在位置
	var queue []int
	//用于存储节点所在位置的深度
	depth := map[int]int{}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 2 {
				queue = append(queue, i*cols+j)
				depth[i*cols+j] = 0
			}
		}
	}

	dr := []int{-1, 0, 1, 0}
	dc := []int{0, -1, 0, 1}

	for len(queue) > 0 {
		q := queue[0]
		queue = queue[1:]
		r, c := q/cols, q%cols

		for k := 0; k < 4; k++ {
			nr, nc := r+dr[k], c+dc[k]
			if nr >= 0 && nr < rows && nc >= 0 && nc < cols && grid[nr][nc] == 1 {
				grid[nr][nc] = 2
				ncode := nr*cols + nc
				queue = append(queue, ncode)
				depth[ncode] = depth[q] + 1
				ans = max(ans, depth[ncode])
			}

		}

		depth[q] = depth[q] + 1
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				return -1
			}
		}
	}

	return ans
}
