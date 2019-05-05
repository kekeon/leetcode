package tree

func searchBST(root *TreeNode1, val int) *TreeNode1 {

	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	}

	left := searchBST(root.Left, val)
	right := searchBST(root.Right, val)

	if left != nil {
		return left
	} else if right != nil {
		return right
	}
	return nil
}
