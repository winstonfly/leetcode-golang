package list

import (
	"leetcode-golang/pkg/top100/entity"
)

func removeNthFromEnd1(head *entity.ListNode, n int) *entity.ListNode {
	l := 0
	current := head
	for {
		if current != nil {
			current = current.Next
			l++
		} else {
			break
		}
	}

	//删除第l -n +1个节点
	//创建一个哨兵节点
	dummy := &entity.ListNode{Next: head}
	index := 1
	prev := dummy
	current = dummy.Next
	for index < l-n+1 {
		prev = current
		current = current.Next
		index++
	}

	//删除
	prev.Next = current.Next

	return dummy.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var resource []*ListNode
	for head != nil {
		tmp := head.Next
		head.Next = nil
		resource = append(resource, head)
		head = tmp
	}

	dummy := &ListNode{}
	current := dummy
	for i := 0; i < len(resource); i++ {
		if i == len(resource)-n {
			continue
		}

		current.Next = resource[i]
		current = current.Next
	}
	return dummy.Next
}
