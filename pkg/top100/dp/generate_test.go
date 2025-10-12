package dp

import "testing"

func TestGenerate(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		v := generate(5)
		t.Log(v)
	})
}
