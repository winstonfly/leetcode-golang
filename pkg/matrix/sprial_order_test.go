package matrix

import "testing"

func TestSprialOrder(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		orders := spiralOrder(matrix)
		t.Log(orders)

	})

	t.Run("test2", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 10, 11, 12},
		}

		orders := spiralOrder(matrix)
		t.Log(orders)
	})
}
