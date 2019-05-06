package tree

func leafSimilar(root1 *TreeNode1, root2 *TreeNode1) bool {
	return intSliceEqual(preOrder(root1), preOrder(root2))
}

func preOrder(root *TreeNode1) []int {
	list := []int{}
	nodeList := []*TreeNode1{}
	nodeList = append(nodeList, root)
	l := len(nodeList)
	for l > 0 {
		node := nodeList[l-1]
		nodeList = nodeList[0 : l-1]

		if node.Left == nil && node.Right == nil {
			list = append(list, node.Val)
		} else {
			if node.Right != nil {
				nodeList = append(nodeList, node.Right)
			}
			if node.Left != nil {
				nodeList = append(nodeList, node.Left)
			}
		}
		l = len(nodeList)
	}

	return list
}

func intSliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	if (a == nil) != (b == nil) {
		return false
	}

	b = b[:len(a)]
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}
