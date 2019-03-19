package tree

import (
	"strings"
	"strconv"
)

var res []string
func binaryTreePaths(root *TreeNode1) []string {

	if root == nil {
		return res
	}

	treePath(root, []string{})

	a := append([]string{}, res...)
	res = []string{}
	return a
}

func treePath(node *TreeNode1, strs []string) {

	if node != nil && node.Left == nil && node.Right == nil {
		strs = append(strs, strconv.Itoa(node.Val))
		res = append(res,strings.Join(strs, "->") )
	}else {
		if node.Left != nil || node.Right != nil  {
			strs = append(strs, strconv.Itoa(node.Val))
			if node.Left != nil {
				treePath(node.Left, strs)
			}

			if node.Right != nil {
				treePath(node.Right, strs)
			}
		}
	}
}
