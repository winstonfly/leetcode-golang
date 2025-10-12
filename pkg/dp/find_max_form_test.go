package dp

import "testing"

func TestFindMaxForm(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		v := findMaxForm([]string{"10", "0001", "111001", "1", "0"}, 5, 3)
		want := 4
		if v != want {
			t.Errorf("findMaxForm() = %d, want %d", v, want)
		}
	})
}
