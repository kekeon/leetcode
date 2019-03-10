package tree

func max(a int, b int) int {
	if a >= b {
		return a
	}

	return b
}

func maxDepth(root *TreeNode) int {

	if root == nil {
		return 0
	}

	return max(maxDepth(root.Right),maxDepth(root.Left)) + 1
}
