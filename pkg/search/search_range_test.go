package search

import "testing"

func TestSearchRange(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{5, 7, 7, 8, 8, 10}
		target := 8
		v := searchRange(nums, target)
		if v[0] != 3 || v[1] != 4 {
			t.Errorf("searchRange failed, target:%v, v:%v", target, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		nums := []int{1}
		target := 0
		v := searchRange(nums, target)
		if v[0] != -1 || v[1] != -1 {
			t.Errorf("searchRange failed, target:%v, v:%v", target, v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		nums := []int{2, 2}
		target := 3
		v := searchRange(nums, target)
		if v[0] != -1 || v[1] != -1 {
			t.Errorf("searchRange failed, target:%v, v:%v", target, v)
		}
	})

	t.Run("test4", func(t *testing.T) {
		nums := []int{3, 3}
		target := 3
		v := searchRange(nums, target)
		if v[0] != 0 || v[1] != 1 {
			t.Errorf("searchRange failed, target:%v, v:%v", target, v)
		}
	})
}
