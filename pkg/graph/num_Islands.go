package graph

// NO.200
func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])

	var ans int
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == '1' {
				infect(grid, i, j)
				ans++
			}
		}
	}

	return ans
}

func infect(grid [][]byte, i, j int) {

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != '1' {
		return
	}

	grid[i][j] = '2'
	infect(grid, i+1, j)
	infect(grid, i-1, j)
	infect(grid, i, j+1)
	infect(grid, i, j-1)
}
