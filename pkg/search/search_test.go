package search

import "testing"

func TestSearch(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := search([]int{4, 5, 6, 7, 0, 1, 2}, 0)
		want := 4
		if v != want {
			t.Errorf("search failed, want:%v, got:%v", want, v)
		}
	})

	t.Run("test1", func(t *testing.T) {
		v := search([]int{4, 5, 6, 7, 8, 0, 1, 2}, 8)
		want := 4
		if v != want {
			t.Errorf("search failed, want:%v, got:%v", want, v)
		}
	})
}
