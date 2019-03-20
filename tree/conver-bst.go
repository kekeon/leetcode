package tree

func convertBST(root *TreeNode1) *TreeNode1 {
	convert(root)
	return root
}

var add int
func convert(root *TreeNode1){
	if root != nil {
		convert(root.Right)
		root.Val += add
		add = root.Val
		convert(root.Left)
	}
}
