package tree

func minDepth(root *TreeNode1) int {

	if root == nil {
		return 0
	}

	l := minDepth(root.Left)
	r := minDepth(root.Right)

	if l == 0 || r == 0 {
		return l + r + 1;
	}

	return min(l, r) + 1
}

func min(a, b int) int {

	if a < b {
		return a
	}

	return b
}
