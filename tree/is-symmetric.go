package tree



func RLsymmetric(p *TreeNode, q *TreeNode) bool {

	if q == nil && p == nil {

		return true
	}

	if q == nil || p == nil {

		return true
	}

	if p.Val != q.Val {
		return false
	}

	return RLsymmetric(p.Left,q.Right) && RLsymmetric(q.Left,p.Right)

}

func isSymmetric(root *TreeNode) bool {

	if root == nil {
		return true
	}

	return RLsymmetric(root.Left, root.Right)
}
