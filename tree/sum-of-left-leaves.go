package tree

func sumOfLeftLeaves(root *TreeNode1) int {

	if root == nil {
		return 0
	}

	return sumLeft(root.Left, 1) + sumLeft(root.Right, 0)
}

func sumLeft(node *TreeNode1, lab int) int {

	if node != nil && node.Left == nil && node.Right == nil {

		if lab == 1 {
			return node.Val
		}
		return 0

	} else if node != nil {
		return sumLeft(node.Left, 1) + sumLeft(node.Right, 0)
	} else {
		return 0
	}
}
