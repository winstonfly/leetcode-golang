package arrays

import "testing"

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	rotate(nums, k)
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("rotate(%v, %d) = %v; want %v", []int{1, 2, 3, 4, 5, 6, 7}, k, nums, expected)
			break
		}
	}

	nums = []int{-1, -100, 3, 99}
	k = 2
	expected = []int{3, 99, -1, -100}
	rotate(nums, k)
	for i := range nums {
		if nums[i] != expected[i] {
			t.Errorf("rotate(%v, %d) = %v; want %v", []int{-1, -100, 3, 99}, k, nums, expected)
			break
		}
	}
}
