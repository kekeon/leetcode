package tree

func rangeSumBST(root *TreeNode1, L int, R int) int {
	if root == nil {
		return 0
	}

	if root.Val > R {
		return rangeSumBST(root.Left, L, R)
	}else if root.Val < L {
		return rangeSumBST(root.Right, L, R)
	}else {
		return root.Val + rangeSumBST(root.Left, L, R) +  rangeSumBST(root.Right, L, R)
	}
}
