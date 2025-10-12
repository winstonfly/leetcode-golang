package list

// NO.160
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var nodes []*ListNode
	for currentA := headA; currentA != nil; currentA = currentA.Next {
		nodes = append(nodes, currentA)
	}

	for currentB := headB; currentB != nil; currentB = currentB.Next {
		for _, node := range nodes {
			if currentB == node {
				return currentB
			}
		}
	}

	return nil
}
