package backtrack

import "testing"

func TestCombinationSum(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := combinationSum([]int{2, 3, 6, 7}, 7)
		expected := [][]int{{2, 2, 3}, {7}}
		if len(v) != len(expected) {
			t.Errorf("Expected %v, but got %v", expected, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := combinationSum([]int{8, 7, 4, 3}, 11)
		expected := [][]int{{8, 3}, {7, 4}, {4, 4, 3}}
		if len(v) != len(expected) {
			t.Errorf("Expected %v, but got %v", expected, v)
		}
	})

	t.Run("test3", func(t *testing.T) {
		v := combinationSum([]int{2, 3, 8, 4}, 6)
		expected := [][]int{{2, 2, 2}, {3, 3}, {2, 4}}
		if len(v) != len(expected) {
			t.Errorf("Expected %v, but got %v", expected, v)
		}
	})
}
