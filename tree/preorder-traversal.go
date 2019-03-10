package tree

import (
	"leetcode/stack"
	"fmt"
)

type TreeNode struct {
	Val   interface{}
	Left  *TreeNode
	Right *TreeNode
}

var arr []interface{}

func visit(v interface{}) {
	arr = append(arr, v)
}

func visitAlongLeftBranch(x *TreeNode, s *stack.Stack) {

	for x != nil {
		// 树节点的val
		visit(x.Val)
		s.Push(x.Right)
		x = x.Left
	}

}

func travPre(x *TreeNode) interface{} {
	stackNode := stack.ConstructorStack()

	for true {
		visitAlongLeftBranch(x, &stackNode)

		if stackNode.Empty() {
			break
		}

		a, ok := stackNode.Pop().(*TreeNode)
		if  !ok {
			fmt.Println("not type *TreeNode")
			break
		}
		x = a
	}

	return arr
}



