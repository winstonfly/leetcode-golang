package list

// NO.206
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		tmp := current.Next
		current.Next = prev
		prev = current
		current = tmp
	}

	return prev
}
