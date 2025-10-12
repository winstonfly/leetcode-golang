package list

func rotateRight(head *ListNode, k int) *ListNode {
	//先获取链表长度
	l := 0
	dummy := &ListNode{Next: head}
	c := dummy.Next
	for c != nil {
		l++
		c = c.Next
	}

	//移动的位置为
	if l == 0 || k%l == 0 {
		return head
	} else {
		position := 0
		if k > l {
			position = k % l
		} else {
			position = k
		}
		prev := dummy
		current := dummy.Next
		index := 1
		for current != nil {
			if index == l-position+1 {
				//找到要翻转的节点起始节点
				tail := current
				for tail.Next != nil {
					tail = tail.Next
				}

				prev.Next = nil
				tail.Next = dummy.Next
				break
			}
			prev = current
			current = current.Next
			index++
		}

		return current
	}
}
