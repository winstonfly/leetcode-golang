package dp

import "testing"

func TestMinPathSum(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
		want := 7
		if v != want {
			t.Errorf("minPathSum() = %d, want %d", v, want)
		}
	})
}
