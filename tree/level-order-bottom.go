package tree

type TreeNode1 struct {
	Val   int
	Left  *TreeNode1
	Right *TreeNode1
}

func visitVal(p []interface{}) []int{

	arr := []int{}
	for _, val := range p {

		if v, ok:= val.(*TreeNode1); ok && v != nil {
			arr = append(arr, v.Val)
		}
	}
	return arr
}

func levelOrderBottom(root *TreeNode1) [][]int {
	q := [][]int{}
	arr := []interface{}{}

	if root == nil {
		return q
	}
	arr = append(arr, root)
	q = append(q, visitVal(arr))

	for len(arr) != 0 {

		a := []interface{}{}

		for _, v := range arr {

			node := v.(*TreeNode1)

			if node.Left != nil {
				a = append(a, node.Left)
			}

			if node.Right != nil {
				a = append(a, node.Right)
			}
		}

		l := visitVal(a)


		arr = a
		if len(q) == 0 {
			q = append([][]int{}, l)

		}else if len(q) > 0 && len(l) > 0{
				c := append([][]int{}, q...)

				q = append([][]int{}, l)

				q = append(q, c...)
		}
	}
	return q

}
