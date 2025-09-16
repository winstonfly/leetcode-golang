package arrays

import "testing"

func TestThreeSum(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []int{-1, 0, 1, 2, -1, -4}
		v := threeSum(tests)
		t.Log(v)
	})

	t.Run("test2", func(t *testing.T) {
		tests := []int{0, 0, 0, 0}
		v := threeSum(tests)
		t.Log(v)
	})

	t.Run("test3", func(t *testing.T) {
		tests := []int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4}
		v := threeSum(tests)
		t.Log(v)
	})

	t.Run("test4", func(t *testing.T) {
		tests := []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}
		v := threeSum(tests)
		t.Log(v)
	})

	t.Run("test5", func(t *testing.T) {
		tests := []int{3, -2, 1, 0}
		v := threeSum(tests)
		t.Log(v)
	})
}
