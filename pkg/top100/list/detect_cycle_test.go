package list

import "testing"

func TestDetectCycle(t *testing.T) {
	t.Run("detectCycle", func(t *testing.T) {
		head := BuildListNode([]int{1, 2, 3, 4, 5})
		p := detectCycle(head)
		if p != nil {
			t.Log(p.Val)
		} else {
			t.Log("No cycle detected")
		}
	})

	t.Run("detectCycle with cycle", func(t *testing.T) {
		head := BuildListNode([]int{1, 2, 3, 4, 5})
		head.Next.Next.Next.Next = head.Next // Create a cycle
		p := detectCycle(head)
		if p != nil {
			t.Log(p.Val)
		} else {
			t.Log("No cycle detected")
		}
	})
}
