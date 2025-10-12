package stacks

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := isValid("()[}]{")
		want := false
		if v != want {
			t.Errorf("isValid failed, want:%v, got:%v", want, v)
		}
	})

	t.Run("test2", func(t *testing.T) {
		v := isValid("[")
		want := false
		if v != want {
			t.Errorf("isValid failed, want:%v, got:%v", want, v)
		}
	})
}
