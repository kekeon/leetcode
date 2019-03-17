package tree

func lowestCommonAncestor(root, p, q *TreeNode1) *TreeNode1 {

	if p != nil && p.Val < root.Val &&  q.Val < root.Val {
		return lowestCommonAncestor(root.Left,p, q)
	}else if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right,p,q)
	}else {
		return root
	}

	return root
}