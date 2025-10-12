package dp

import "testing"

func TestMaxProduct(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []int{1, 2, 3}
		v := maxProduct(tests)
		if v != 6 {
			t.Errorf("maxProduct(%d) = %d, want 6", tests, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		tests := []int{0, -1, 4}
		v := maxProduct(tests)
		if v != 4 {
			t.Errorf("maxProduct(%d) = %d, want 4", tests, v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		tests := []int{2, 3, -2, 4}
		v := maxProduct(tests)
		if v != 6 {
			t.Errorf("maxProduct(%d) = %d, want 6", tests, v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		tests := []int{-1, -2, -9, -6}
		v := maxProduct(tests)
		if v != 108 {
			t.Errorf("maxProduct(%d) = %d, want 108", tests, v)
		}
	})
}
