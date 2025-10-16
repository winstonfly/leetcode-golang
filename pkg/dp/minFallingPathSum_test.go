package dp

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := minFallingPathSum([][]int{
			{2, 1, 3},
			{6, 5, 4},
			{7, 8, 9},
		})
		want := 13
		if v != want {
			t.Errorf("minFallingPathSum() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := minFallingPathSum([][]int{
			{-19, 57},
			{-40, -5},
		})
		want := -59
		if v != want {
			t.Errorf("minFallingPathSum() = %d, want %d", v, want)
		}
	})
}
