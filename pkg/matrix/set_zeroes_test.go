package matrix

import "testing"

func TestSetZeroes(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		matrix := [][]int{
			{1, 1, 1},
			{1, 0, 1},
			{1, 1, 1},
		}
		expected := [][]int{
			{1, 0, 1},
			{0, 0, 0},
			{1, 0, 1},
		}
		setZeroes(matrix)
		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] != expected[i][j] {
					t.Errorf("setZeroes() = %v; want %v", matrix, expected)
					return
				}
			}
		}
	})

	t.Run("test2", func(t *testing.T) {
		matrix := [][]int{
			{0, 1, 2, 0},
			{3, 4, 5, 2},
			{1, 3, 1, 5},
		}
		expected := [][]int{
			{0, 0, 0, 0},
			{0, 4, 5, 0},
			{0, 3, 1, 0},
		}
		setZeroes(matrix)
		for i := range matrix {
			for j := range matrix[i] {
				if matrix[i][j] != expected[i][j] {
					t.Errorf("setZeroes() = %v; want %v", matrix, expected)
					return
				}
			}
		}
	})

}
