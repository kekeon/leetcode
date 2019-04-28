package tree

import (
	"strconv"
)

func tree2str(n *TreeNode1) string {

	if n == nil {
		return ""
	}

	if n.Left == nil && n.Right == nil {
		return strconv.Itoa(n.Val)
	}

	if n.Left != nil && n.Right == nil {
		return  strconv.Itoa(n.Val) + "(" + tree2str(n.Left) +")"
	}else if n.Left == nil && n.Right != nil {

		return strconv.Itoa(n.Val) + "()"+ "(" + tree2str(n.Right) +")"
	}

	return strconv.Itoa(n.Val) + "(" + tree2str(n.Left) +")" + "(" + tree2str(n.Right) +")"

}
