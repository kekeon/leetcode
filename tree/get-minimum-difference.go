package tree

import "math"

func getMinimumDifference(root *TreeNode1) int {

	minLeft := math.MaxInt64
	minRight := math.MaxInt64

	if root.Left != nil {
		node := root.Left
		for node.Right != nil {
			node = node.Right
		}

		minLeft = min(abs(root.Val-node.Val), getMinimumDifference(root.Left))
	}

	if root.Right != nil {
		node := root.Right
		for node.Left != nil {
			node = node.Left
		}

		minRight = min(abs(root.Val-node.Val), getMinimumDifference(root.Right))
	}

	return min(minLeft, minRight)
}