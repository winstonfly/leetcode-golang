package search

import "testing"

func TestSearchMatrix(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		matrix := [][]int{
			[]int{1, 3, 5, 7},
			[]int{10, 11, 16, 20},
			[]int{23, 30, 34, 60},
		}

		target := 3
		v := searchMatrix(matrix, target)
		if !v {
			t.Errorf("searchMatrix failed, target:%v, v:%v", target, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		matrix := [][]int{
			[]int{1},
		}
		target := 1
		v := searchMatrix(matrix, target)
		if !v {
			t.Errorf("searchMatrix failed, target:%v, v:%v", target, v)
		}
	})
}
