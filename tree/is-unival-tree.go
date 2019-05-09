package tree

func isUnivalTree(root *TreeNode1) bool {
	return univalTree(root, root.Val)
}


func univalTree(node *TreeNode1, v int ) bool {

	if node == nil {
		return true
	}

	if node.Val != v {
		return false
	}else {
		return  univalTree(node.Right, node.Val) && univalTree(node.Left, node.Val)
	}

}
