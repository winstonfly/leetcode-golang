package list

import "testing"

func TestMiddleNode(t *testing.T) {
	t.Run("middleNode", func(t *testing.T) {
		head := BuildListNode([]int{1, 2, 3, 4, 5})
		p := middleNode(head)
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	})
}
