package tree




func isBalanced(root *TreeNode1) bool {

	if (root == nil) {
		return true
	}

	left := nodeHeight(root.Left)
	right := nodeHeight(root.Right)

	if (left > right && left - right > 1) || (left < right && right - left > 1) {
		return false
	}

	if isBalanced(root.Left) && isBalanced(root.Right) {
		return true
	}

	return false
}

func nodeHeight(node *TreeNode1) int {

	if node != nil {
		return max(nodeHeight(node.Right), nodeHeight(node.Left)) + 1
	}

	return 0
}