package tree

func findSecondMinimumValue(root *TreeNode1) int {

	if root == nil {
		return -1
	}

	if root.Left == nil && root.Right == nil {
		return -1
	}

	l := root.Left.Val
	r := root.Right.Val

	if l == root.Val {
		l = findSecondMinimumValue(root.Left)
	}

	if r == root.Val {
		r = findSecondMinimumValue(root.Right)
	}

	if l != -1 && r != -1 {
		return min(l, r)
	}

	if r == -1 {
		return l
	}

	if l == -1 {
		return r
	}

	return -1
}
