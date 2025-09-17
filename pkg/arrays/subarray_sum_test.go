package arrays

import "testing"

func TestSubarraySum(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		nums := []int{1, 1, 1}
		k := 2
		expected := 2
		if result := subarraySum(nums, k); result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("test2", func(t *testing.T) {
		nums := []int{1, 2, 3}
		k := 3
		expected := 2
		if result := subarraySum(nums, k); result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("test3", func(t *testing.T) {
		nums := []int{-1, -1, 1}
		k := 0
		expected := 1
		if result := subarraySum(nums, k); result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})

	t.Run("test4", func(t *testing.T) {
		nums := []int{1, 2, 1, 2, 1}
		k := 3
		expected := 4
		if result := subarraySum(nums, k); result != expected {
			t.Errorf("expected %d, got %d", expected, result)
		}
	})
}
