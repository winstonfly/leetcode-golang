package list

import "testing"

func TestSortList(t *testing.T) {
	t.Run("sortList", func(t *testing.T) {
		tests := BuildListNode([]int{-84, 142, 41, -17, -71, 170, 186, 183, -21, -76, 76, 10, 29, 81, 112, -39, -6, -43, 58, 41, 111, 33, 69, 97, -38, 82, -44, -7, 99, 135, 42, 150, 149, -21, -30, 164, 153, 92, 180, -61, 99, -81, 147, 109, 34, 98, 14, 178, 105, 5, 43, 46, 40, -37, 23, 16, 123, -53, 34, 192, -73, 94, 39, 96, 115, 88, -31, -96, 106, 131, 64, 189, -91, -34, -56, -22, 105, 104, 22, -31, -43, 90, 96, 65, -85, 184, 85, 90, 118, 152, -31, 161, 22, 104, -85, 160, 120, -31, 144, 115})
		p := sortList(tests)
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	})

	t.Run("sortList2", func(t *testing.T) {
		tests := BuildListNode([]int{4, 2, 1, 3})
		p := sortList(tests)
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	})

}
