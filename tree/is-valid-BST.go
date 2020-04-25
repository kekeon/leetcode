package tree

import "math"


func isValidBSTNumber(root *TreeNode1, min, max int) bool {

	if root == nil {
		return true
	}

	if root.Val <= min  || root.Val >= max{
		return false
	}


	return isValidBSTNumber(root.Left, min, root.Val) && isValidBSTNumber(root.Right, root.Val, max)

}

func isValidBST(root *TreeNode1) bool {

	return isValidBSTNumber(root, math.MinInt64, math.MaxInt64)
}
