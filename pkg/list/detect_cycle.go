package list

// NO.142
func detectCycle(head *ListNode) *ListNode {

	m := make(map[*ListNode]bool, 1)
	for head != nil {
		if m[head] {
			return head
		}

		m[head] = true
		head = head.Next
	}

	return nil
}

func detectCycle2(head *ListNode) *ListNode {
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}

	return slow
}
