package tree

func hasPathSum(root *TreeNode1, sum int) bool {

	if root == nil {
		return false
	}

	if root.Left == nil && root.Right == nil && root.Val == sum {
		return true
	}

	return hasPathSum(root.Left, sum - root.Val) || hasPathSum(root.Right, sum - root.Val)

}
