package backtrack

import (
	"testing"
)

func TestPermute(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		tests := []int{1, 2, 3}
		want := [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}
		//
		t.Log(permute(tests), want)
	})
}
