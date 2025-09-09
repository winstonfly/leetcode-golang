package stacks

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []int{2, 1, 5, 6, 2, 3}
		v := largestRectangleArea(tests)
		if v != 10 {
			t.Errorf("got %d, want 10", v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		tests := []int{2, 0, 2}
		v := largestRectangleArea(tests)
		if v != 2 {
			t.Errorf("got %d, want 2", v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		tests := []int{0, 9}
		v := largestRectangleArea(tests)
		if v != 9 {
			t.Errorf("got %d, want 9", v)
		}
	})

	t.Run("test4", func(t *testing.T) {
		tests := []int{4, 2, 0, 3, 2, 4, 3, 4}
		v := largestRectangleArea(tests)
		if v != 10 {
			t.Errorf("got %d, want 10", v)
		}
	})
}
