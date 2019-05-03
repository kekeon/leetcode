package tree

var list []int
func findTarget(root *TreeNode1, k int) bool {
	list = []int{}
	inorder(root)

	start := 0
	last := len(list) - 1
	for start < last {
		sum := list[start] + list[last]

		if sum < k {
			start += 1
		}else if sum > k {
			last -=1
		}else {
			return true
		}
	}

	return false
}
func inorder(node *TreeNode1) {
	if node != nil {
		inorder(node.Left)
		list = append( list, node.Val)
		inorder(node.Right)
	}
}
