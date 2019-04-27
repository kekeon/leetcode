package tree


func findTilt(root *TreeNode1) int {
	if root == nil {
		return 0
	}

	res := 0
	nodeSum(root, &res)
	return res
}

func nodeSum(node *TreeNode1, res *int) int {
	if node == nil {
		*res += 0
		return 0
	}

	l := nodeSum(node.Left, res)
	r := nodeSum(node.Right, res)

	*res += intAbs(l - r)

	return l + r + node.Val
}

func intAbs(i int) int {

	if i >= 0 {
		return 0
	}

	return -i
}
