package tree

import (
	"leetcode/stack"
)

// 层次遍历
func LevelTraversal(x *TreeNode) []interface{}{

	q := stack.ConstructorQueue()
	arr := []interface{}{}

	q.Push(x)

	for !q.Empty() {

		node := q.Pop().(*TreeNode)

		arr = append(arr, node.Val)

		if node.Left != nil {
			q.Push(node.Left)
		}

		if node.Right != nil {
			q.Push(node.Right)
		}
	}

	return arr

}