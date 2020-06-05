package tree

func levelOrder(root *TreeNode1) [][]int {

	res := [][]int{}
	if root == nil {
		return res
	}

	// 层级队列
	queue := []TreeNode1{}
	queue = append(queue, *root)

	// 循环队列
	for len(queue) > 0 {
		lev := []int{}
		// 缓存新一一层队列
		levelNode := []TreeNode1{}

		// 遍历父级队列
		for _, node := range queue {

			lev = append(lev, node.Val)


			if node.Left != nil {
				levelNode = append(levelNode, *node.Left)
			}

			if node.Right != nil {
				levelNode = append(levelNode, *node.Right)
			}
		}
		// 更新队列
		queue = levelNode
		res = append(res, lev)
	}

	return res
}
