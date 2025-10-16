package dp

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		grid := [][]int{
			{0, 0, 0},
			{0, 1, 0},
			{0, 0, 0},
		}
		v := uniquePathsWithObstacles(grid)
		want := 2
		if v != want {
			t.Errorf("uniquePathsWithObstacles() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		grid := [][]int{
			{0, 1},
			{0, 0},
		}
		v := uniquePathsWithObstacles(grid)
		want := 1
		if v != want {
			t.Errorf("uniquePathsWithObstacles() = %d, want %d", v, want)
		}
	})
}
