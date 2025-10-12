package dp

import "testing"

func TestChange(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := change(5, []int{1, 2, 5})
		want := 4
		if v != want {
			t.Errorf("change() = %d, want %d", v, want)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := change(3, []int{2})
		want := 0
		if v != want {
			t.Errorf("change() = %d, want %d", v, want)
		}
	})

	t.Run("test3", func(t *testing.T) {
		v := change(10, []int{10})
		want := 1
		if v != want {
			t.Errorf("change() = %d, want %d", v, want)
		}
	})
}
