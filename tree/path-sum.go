package tree

func pathSum(root *TreeNode1, sum int) int {

	if root == nil {
		return 0
	}

	return pathSum(root.Left, sum ) + pathSum(root.Right, sum) + addPath(root, sum,0 )
}

func addPath(node *TreeNode1, sum int, count int) int{

	if node == nil {
		return 0
	}else {

		if node.Val == sum {
			return 1
		}

		return addPath(node.Left, sum - node.Val, count) + addPath(node.Right, sum - node.Val, count)
	}

}