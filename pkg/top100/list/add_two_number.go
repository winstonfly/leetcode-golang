package list

func addTwoNumbers1(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil && l2 != nil {
		curr := &ListNode{}
		curr.Val = (l1.Val + l2.Val + carry) % 10
		current.Next = curr
		current = curr
		if l1.Val+l2.Val+carry >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		l1.Val += carry
		curr := &ListNode{
			Val: l1.Val % 10,
		}
		current.Next = curr
		current = curr
		if l1.Val >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		l1 = l1.Next
	}

	for l2 != nil {
		l2.Val += carry
		if l2.Val >= 10 {
			carry = 1
		} else {
			carry = 0
		}
		curr := &ListNode{
			Val: l2.Val % 10,
		}
		current.Next = curr
		current = curr
		l2 = l2.Next
	}
	if carry == 1 {
		current.Next = &ListNode{Val: 1}
	}

	return dummy.Next
}

func addTwoNumbers(l1, l2 *ListNode) *ListNode {
	var result []int

	for l1 != nil && l2 != nil {
		result = append(result, l1.Val+l2.Val)
		l1 = l1.Next
		l2 = l2.Next
	}

	for l1 != nil {
		result = append(result, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		result = append(result, l2.Val)
		l2 = l2.Next
	}

	dummy := &ListNode{}
	current := dummy
	carry := 0
	for i, v := range result {
		if carry != 0 {
			v += carry
		}
		if v >= 10 {
			carry = v / 10
			current.Next = &ListNode{Val: v - 10}
		} else {
			carry = 0
			current.Next = &ListNode{Val: v}
		}

		if i == len(result)-1 {
			if carry != 0 {
				current.Next.Next = &ListNode{Val: carry}
			}
		}

		current = current.Next

	}

	return dummy.Next
}
