package search

import "testing"

func TestInsert(t *testing.T) {
	t.Run("test1", func(t *testing.T) {

		tests := []int{1, 3, 5, 6}
		target := 5
		v := searchInsert(tests, target)
		if v != 2 {
			t.Errorf("searchInsert(%v, %v) = %d, want %d", tests, target, v, 2)
		}
	})

	t.Run("test2", func(t *testing.T) {
		tests := []int{1, 3, 5, 6}
		target := 2
		v := searchInsert(tests, target)
		if v != 1 {
			t.Errorf("searchInsert(%v, %v) = %d, want %d", tests, target, v, 1)
		}
	})

	t.Run("test3", func(t *testing.T) {

		tests := []int{1, 3, 5, 6}
		target := 7
		v := searchInsert(tests, target)
		if v != 4 {
			t.Errorf("searchInsert(%v, %v) = %d, want %d", tests, target, v, 4)
		}
	})

	t.Run("test4", func(t *testing.T) {

		tests := []int{1, 3, 5, 6}
		target := 0
		v := searchInsert(tests, target)
		if v != 0 {
			t.Errorf("searchInsert(%v, %v) = %d, want %d", tests, target, v, 0)
		}
	})
}
