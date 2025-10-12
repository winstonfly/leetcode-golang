package backtrack

import (
	"testing"
)

func TestSubsets(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []int{1, 2, 3}
		want := [][]int{{}, {1}, {1, 2}, {1, 2, 3}, {1, 3}, {2}, {2, 3}, {3}}
		got := subsets(tests)
		t.Log(got, want)
	})
}
