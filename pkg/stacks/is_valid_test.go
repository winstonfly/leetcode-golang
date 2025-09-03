package stacks

import "testing"

func TestIsValid(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := isValid("()[}]{")
		t.Log(v)
	})

	t.Run("test2", func(t *testing.T) {
		v := isValid("[")
		t.Log(v)
	})
}
