package tree

func averageOfLevels(root *TreeNode) []float64 {

}

func average(l, r *TreeNode1) float64 {

	if l != nil && r != nil {
		return (average(l.Right, l.Left) + average(r.Right, r.Left)) / 2
	}else if l != nil && r == nil {
		return average(l.Right, l.Left)
	}else if l == nil && r != nil {
		return average(r.Right, r.Left)
	}
}
