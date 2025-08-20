package arrays

import (
	"reflect"
	"testing"
)

func TestMoveZeros(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	expected := []int{1, 3, 12, 0, 0}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected %v, got %v", expected, nums)
	}

	nums = []int{0}
	moveZeroes(nums)
	expected = []int{0}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected %v, got %v", expected, nums)
	}

	nums = []int{1, 2, 3}
	moveZeroes(nums)
	expected = []int{1, 2, 3}
	if !reflect.DeepEqual(nums, expected) {
		t.Errorf("expected %v, got %v", expected, nums)
	}
}
