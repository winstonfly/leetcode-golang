package hp

import "testing"

func TestMedianFinder(t *testing.T) {

	t.Run("TestMedianFinder", func(t *testing.T) {
		finder := Constructor()
		finder.AddNum(1)
		finder.AddNum(2)
		v := finder.FindMedian()
		if v != 1.5 {
			t.Errorf("expected 1.5, got %v", v)
		}
		finder.AddNum(3)
		v = finder.FindMedian()
		if v != 2 {
			t.Errorf("expected 2, got %v", v)
		}
	})

}
