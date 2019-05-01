package tree

func convertBST(root *TreeNode1) *TreeNode1 {
	convert(root, 0)
	return root
}

func convert(root *TreeNode1, sum int) int{

	if root == nil {
		return sum
	}

	s := convert(root.Right, sum)

	root.Val += s

	s = root.Val

	s = convert(root.Left , s)

	return s
}
