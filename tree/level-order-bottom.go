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
	arr = append(arr, root)
	q = append(q, visitVal(arr))

	for len(arr) != 0 {

		a := []interface{}{}

		for _, v := range arr {

			node := v.(*TreeNode1)
			
			if node.Left != nil {
				a = append(arr, node.Left)
			}

			if node.Right != nil {
				a = append(arr, node.Right)
			}
		}
		
		arr = a

		
		
		l := visitVal(a)

		if len(q) > 0 {

			c := append([][]int{}, q...)

			q = append([][]int{},l)

			q = append([][]int{},c...)
		}else {
			q = append([][]int{}, l)
		}


	}

	return q

}
