package list

// NO.141
func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for slow != nil && slow.Next != nil && fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}

		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
