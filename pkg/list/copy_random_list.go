package list

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheNodes map[*Node]*Node

// NO.138
func copyRandomList(head *Node) *Node {
	cacheNodes = make(map[*Node]*Node)
	return deepCopy(head)
}

// NO.138
func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if newNode, ok := cacheNodes[node]; ok {
		return newNode
	}
	newNode := &Node{Val: node.Val}
	cacheNodes[node] = newNode
	//在回溯的过程中，构建 next 和 random 指针
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}
