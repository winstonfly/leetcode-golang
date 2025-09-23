package list

import "testing"

func TestCopyRandomList(t *testing.T) {
	t.Run("test1", func(t *testing.T) {
		head := &Node{Val: 7}
		head.Next = &Node{Val: 13}
		head.Next.Next = &Node{Val: 11}
		head.Next.Next.Next = &Node{Val: 10}
		head.Next.Next.Next.Next = &Node{Val: 1}

		head.Random = nil
		head.Next.Random = head
		head.Next.Next.Random = head.Next.Next.Next.Next
		head.Next.Next.Next.Random = head.Next.Next
		head.Next.Next.Next.Next.Random = head

		t.Log(copyRandomList(head))
	})

	t.Run("test2", func(t *testing.T) {
		head := &Node{Val: 1}
		head.Next = &Node{Val: 2}

		head.Random = head.Next
		head.Next.Random = head.Next

		t.Log(copyRandomList(head))
	})

	t.Run("test3", func(t *testing.T) {
		var head *Node

		t.Log(copyRandomList(head))
	})
}
