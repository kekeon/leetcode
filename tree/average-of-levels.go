package tree

func averageOfLevels(root *TreeNode1) []float64 {

	res := []float64{}

	arr := []*TreeNode1{}

	if root == nil {
		return res
	}
	res = append(res, float64(root.Val))
	arr = append(arr, root)

	for len(arr) != 0 {

		a := []*TreeNode1{}

		for _, v := range arr{
			if v.Left != nil {
				a = append(a, v.Left)
			}

			if v.Right != nil {
				a = append(a, v.Right)
			}
		}

		if len(a) != 0 {
			e := averageVal(a)
			res  = append(res, e)
		}
		arr = a

	}
	return res
}

func averageVal(p []*TreeNode1) float64{
	sum := 0
	for _, e := range p {
		sum += e.Val
	}
	return float64(sum)/float64(len(p))
}
