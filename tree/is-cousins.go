package tree

func isCousins(root *TreeNode, x int, y int) bool {
	if root.Val == x || root.Val == y {
		return false
	}
	dx, fx := depth(root, x, nil)
	dy, fy := depth(root, y, nil)
	if dx == dy && fx.Val != fy.Val {
		return true
	} else {
		return false
	}
}

func depth(root *TreeNode, x int, father *TreeNode) (int, *TreeNode) {
	if root == nil {
		return 0, nil
	}
	if root.Val == x {
		return 1, father
	}
	var f *TreeNode
	left, fl := depth(root.Left, x, root)
	right, fr := depth(root.Right, x, root)
	if fl != nil {
		f = fl
	}
	if fr != nil {
		f = fr
	}
	if left != 0 {
		return left + 1, f
	}
	if right != 0 {
		return right + 1, f
	}
	return 0, nil
}