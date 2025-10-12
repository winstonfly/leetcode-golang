package list

import "testing"

func TestHasCycle(t *testing.T) {

	t.Run("test1", func(t *testing.T) {
		head := &ListNode{Val: 3}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = &ListNode{Val: 0}
		head.Next.Next.Next = &ListNode{Val: -4}
		head.Next.Next.Next.Next = head.Next

		t.Log(hasCycle(head))
	})

	t.Run("test2", func(t *testing.T) {
		head := &ListNode{Val: 1}
		head.Next = &ListNode{Val: 2}
		head.Next.Next = head

		t.Log(hasCycle(head))
	})

	t.Run("test3", func(t *testing.T) {
		head := &ListNode{Val: 1}

		t.Log(hasCycle(head))
	})
}
