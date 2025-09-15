package list

type Node struct {
	Val    int
	Node   *Node
	Random *Node
}

var cachedNodes map[*Node]*Node

// NO.138
func copyRandomList(head *Node) *Node {
	cachedNodes = make(map[*Node]*Node, 1)
	return deepCopy(head)
}

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}

	if n, ok := cachedNodes[node]; ok {
		return n
	}

	newNode := &Node{}
	newNode.Val = node.Val
	newNode.Node = deepCopy(node.Node)
	newNode.Random = deepCopy(node.Random)

	cachedNodes[newNode] = newNode
	return newNode
}
