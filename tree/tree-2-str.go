package tree

import (
	"strconv"
)

func tree2str(t *TreeNode1) string {

	if t == nil {
		return ""
	}

	// 1. 总体思路：t.val + (t.Left) + (t.Right)

	if t.Left == nil && t.Right == nil {

		return strconv.Itoa(t.Val)
	}

	// 2. 处理特殊case, 右子树为nil, 则右() 可省略
	if t.Left != nil && t.Right == nil {

		return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")"

	// 3. 左子树为nil, 则左()不 可省略
	} else if t.Left == nil && t.Right != nil {

		return strconv.Itoa(t.Val) + "()" + "(" + tree2str(t.Right) + ")"
	}

	return strconv.Itoa(t.Val) + "(" + tree2str(t.Left) + ")" + "(" + tree2str(t.Right) + ")"

}
