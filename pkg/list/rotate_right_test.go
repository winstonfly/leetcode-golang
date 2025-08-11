package list

import "testing"

func TestRotateRight(t *testing.T) {
	t.Run("rotateRight", func(t *testing.T) {
		head := BuildListNode([]int{1, 2, 3, 4, 5})
		k := 3
		p := rotateRight(head, k)
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	})

	t.Run("rotateRight2", func(t *testing.T) {
		head := BuildListNode([]int{0, 1, 2})
		k := 4
		p := rotateRight(head, k)
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	})

	t.Run("rotateRight3", func(t *testing.T) {
		head := BuildListNode([]int{1, 2})
		k := 5
		p := rotateRight(head, k)
		for p != nil {
			t.Log(p.Val)
			p = p.Next
		}
	})

}
