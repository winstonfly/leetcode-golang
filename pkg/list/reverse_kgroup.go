package list

// NO.25
func reverseKGroup(head *ListNode, k int) *ListNode {
	dumm := &ListNode{Next: head}
	current := dumm.Next
	prev := dumm
	index := 1

	for current != nil {
		if index%k == 0 {
			tHead := prev
			curr := tHead
			var next *ListNode
			for i := 0; i < k; i++ {
				curr = curr.Next
				if i == k-1 {
					next = curr.Next
					curr.Next = nil
				}
			}
			tmpHead, tmpTail := rotate(tHead.Next)
			prev.Next = tmpHead
			tmpTail.Next = next
			prev = tmpTail
			current = next
		} else {
			current = current.Next
		}

		index++
	}

	return dumm.Next
}

func rotate(head *ListNode) (*ListNode, *ListNode) {

	dummy := &ListNode{Val: -1, Next: head}
	current := dummy.Next
	prev := dummy

	var tmpHead, tmpTail *ListNode
	for current != nil {
		prev.Next = current.Next
		if tmpTail == nil {
			tmpTail = current
			tmpHead = tmpTail
			tmpTail.Next = nil
		} else {
			current.Next = tmpTail
			tmpTail = current
		}
		current = prev.Next
	}

	return tmpTail, tmpHead
}
