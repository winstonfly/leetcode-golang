package tree

import "testing"

func TestMedianFinder(t *testing.T) {

	t.Run("TestMedianFinder", func(t *testing.T) {
		finder := Constructor()
		finder.AddNum(1)
		finder.AddNum(2)
		v := finder.FindMedian()
		finder.AddNum(3)
		v = finder.FindMedian()
		t.Log(v)
	})

}
