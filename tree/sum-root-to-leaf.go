package tree

var sum int
func sumRootToLeaf(root *TreeNode1) int {
	return dfs(root, 0)
}

func dfs(node *TreeNode1, v int) int {
	v = v << 1 | node.Val
	if node.Left == nil && node.Right == nil {
		sum  += v
		return sum
	}

	if node.Left != nil {
		dfs(node.Left, v)
	}

	if node.Right != nil {
		dfs(node.Right, v)
	}
	return sum
}
