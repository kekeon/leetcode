package tree

func trimBST(root *TreeNode1, L int, R int) *TreeNode1 {

	if root == nil {
		return nil
	}

	if root.Val > R {
		return trimBST(root.Right, L, R)
	}else if root.Val < L {
		return trimBST(root.Left, L, R)
	}else {
		root.Left = trimBST(root.Left, L, R)
		root.Right = trimBST(root.Right, L, R)
		return root
	}
}
