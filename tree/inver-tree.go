package tree

func invertTree(root *TreeNode1) *TreeNode1 {

	if root == nil {
		return nil
	}

	c := root.Left
	root.Left = invertTree(root.Right)
	root.Right = invertTree(c)

	return root
}
