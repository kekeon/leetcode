package tree


var maxPathLength int
func longestUnivaluePath(root *TreeNode1) int {
	maxPathLength = 0
	if root == nil {
		return 0
	}

	maxLength(root, root.Val)

	return maxPathLength
}

func maxLength(node *TreeNode1, val int) int{

	if node == nil {
		return 0
	}

	left := maxLength(node.Left, node.Val)

	right := maxLength(node.Right, node.Val)

	maxPathLength = max(maxPathLength, left + right)

	if node.Val == val {
		return max(left, right) + 1
	}
	return 0

}
