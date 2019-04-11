package tree

func diameterOfBinaryTree(root *TreeNode1) int {
	return maxPath(root)[1]
}

func maxPath(node *TreeNode1) []int {

	if node == nil {
		return []int{0, 0}
	}

	left := maxPath(node.Left)
	right := maxPath(node.Right)

	total := left[0] + right[0]

	max := maxVal(left[1], right[1])


	return []int{maxVal(left[0] + 1, right[0] + 1), maxVal(max, total)}
}

func maxVal(a,b int) int {
	if a > b {
		return a
	}
	return b
}
