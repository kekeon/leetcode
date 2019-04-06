package tree

func findMode(root *TreeNode1) []int {

	nodeMap := toMap(root, map[int]int{})

	var mv int
	var res []int

	for k,v := range nodeMap {

		if v > mv {
			mv = v
			res = []int{}
			res = append(res,k)
		}else if v == mv {
			res = append(res,k)
		}
	}

	return res
}

func toMap(node *TreeNode1, m map[int]int) map[int]int {

	if node != nil {
		m[node.Val]++
		toMap(node.Right, m)
		toMap(node.Left, m)
	}
	return m
}
