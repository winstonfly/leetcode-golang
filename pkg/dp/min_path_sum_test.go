package dp

import "testing"

func TestMinPathSum(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := minPathSum([][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}})
		t.Log(v)
	})
}
