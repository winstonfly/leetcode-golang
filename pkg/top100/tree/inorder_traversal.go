package tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Pair struct {
	Node  *TreeNode
	Color int // 0表示白色，1表示灰色
}

func BuildTreeNode(arr []int) *TreeNode {
	if len(arr) == 0 || arr[0] == -1 {
		return nil
	}

	root := &TreeNode{Val: arr[0]}
	queue := []*TreeNode{root}
	index := 1

	for index < len(arr) {
		node := queue[0]
		queue = queue[1:]

		// 处理左子节点
		if index < len(arr) {
			if arr[index] != -1 {
				node.Left = &TreeNode{Val: arr[index]}
				queue = append(queue, node.Left)
			}
			index++
		}

		// 处理右子节点
		if index < len(arr) {
			if arr[index] != -1 {
				node.Right = &TreeNode{Val: arr[index]}
				queue = append(queue, node.Right)
			}
			index++
		}
	}

	return root
}

// 中序遍历二叉树， 从根节点开始，先遍历左子树，再访问根节点，最后遍历右子树
func inorderTraversal(root *TreeNode) []int {
	var ans []int

	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		ans = append(ans, node.Val)
		inorder(node.Right)
	}

	inorder(root)

	return ans
}

func inorderTraversal2(root *TreeNode) []int {
	var stack []Pair
	var ans []int
	stack = append(stack, Pair{Node: root, Color: 0})
	for len(stack) > 0 {
		p := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if p.Color == 0 {
			if p.Node != nil {
				stack = append(stack, Pair{Node: p.Node.Right, Color: 0})
				stack = append(stack, Pair{Node: p.Node, Color: 1})
				stack = append(stack, Pair{Node: p.Node.Left, Color: 0})
			}
		} else {
			if p.Node != nil {
				ans = append(ans, p.Node.Val)
			}
		}
	}
	return ans
}

// 后序遍历二叉树， 从根节点开始，先遍历左子树，再遍历右子树，最后访问根节点
func inbackTraversal(root *TreeNode) []int {
	var ans []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		inorder(node.Right)
		ans = append(ans, node.Val) // 这里是后序遍历，所以访问根节点在最后
	}

	inorder(root)

	return ans
}

func infrontTraversal(root *TreeNode) []int {
	var ans []int
	var inorder func(node *TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		ans = append(ans, node.Val) // 这里是前序遍历，所以访问根节点在最前面
		inorder(node.Left)
		inorder(node.Right)
	}

	inorder(root)

	return ans
}

func levelOrder(root *TreeNode) [][]int {
	var ans [][]int
	start := []*TreeNode{root}
	for len(start) > 0 {
		var queue []*TreeNode
		var tmp []int
		for _, node := range start {
			if node == nil {
				continue
			}
			tmp = append(tmp, node.Val)
			queue = append(queue, node.Left, node.Right)
		}
		if len(tmp) != 0 {
			ans = append(ans, tmp)
		}
		start = queue
	}

	return ans
}
