package list

func middleNode(head *ListNode) *ListNode {

	dummy := &ListNode{Next: head}
	c := dummy.Next
	l := 1
	for c.Next != nil {
		c = c.Next
		l++
	}
	position := l/2 + 1
	index := 1
	current := dummy.Next
	for current != nil && index < position {
		current = current.Next
		index++
	}

	return current

}
