package matrix

import "testing"

func TestSearchMatrix(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		matrix := [][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}

		target := 8
		result := searchMatrix(matrix, target)
		t.Log(result)
	})

	t.Run("test2", func(t *testing.T) {
		matrix := [][]int{
			{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30},
		}

		target := 5
		result := searchMatrix(matrix, target)
		t.Log(result)
	})

	t.Run("test3", func(t *testing.T) {
		matrix := [][]int{
			{-5},
		}

		target := -5
		result := searchMatrix(matrix, target)
		t.Log(result)
	})
}
