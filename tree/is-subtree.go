package tree

func isSubtree(s *TreeNode1, t *TreeNode1) bool {
	if s == nil && t == nil {
		return true
	}

	if s == nil && t != nil {
		return false
	}

	return sameTree(s, t) || s != nil && t != nil && isSubtree(s.Left, t.Left) || s != nil && t != nil && isSubtree(s.Right, t.Right)
}

func sameTree(s, t *TreeNode1) bool {

	if s == nil && t == nil {
		return true
	}

	if s == nil || t == nil {
		return false
	}

	return s != nil && t != nil && s.Val == t.Val && sameTree(s.Left, t.Left) && sameTree(s.Right, t.Right)

}
