package matrix

import "testing"

func TestRotate(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		rotate(matrix)

		t.Log(matrix)
	})
}
