package tree

// NO.437
func pathSum1(root *TreeNode, sum int) int {
	var ans int

	if root == nil {
		return 0
	}

	var rootSum func(root *TreeNode, sum int)
	rootSum = func(root *TreeNode, sum int) {
		if root == nil {
			return
		}

		if root.Val == sum {
			ans++
		}

		rootSum(root.Left, sum-root.Val)
		rootSum(root.Right, sum-root.Val)
	}

	rootSum(root, sum)

	ans += pathSum(root.Left, sum)
	ans += pathSum(root.Right, sum)
	return ans
}

func pathSum(root *TreeNode, sum int) int {
	var ans int
	prefixSum := make(map[int]int) //用来保存已经计算过的结点和的个数
	prefixSum[0] = 1
	var dfs func(node *TreeNode, currSum int)
	dfs = func(node *TreeNode, currSum int) {
		if node == nil {
			return
		}

		currSum += node.Val
		if count, ok := prefixSum[currSum-sum]; ok {
			ans += count
		}
		prefixSum[currSum]++
		dfs(node.Left, currSum)
		dfs(node.Right, currSum)
		prefixSum[currSum]--

	}

	return ans
}
