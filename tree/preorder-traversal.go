package tree

import (
	"leetcode/stack"
	"fmt"
)


// 先序遍历的实现
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

	for {
		// 树节点的val

		if(x != nil ) {
			visit(x.Val)
			s.Push(x.Right)
			x = x.Left
		}else {
			visit(nil)
			break
		}
	}

}

func travPre(x *TreeNode) interface{} {
	stackNode := stack.ConstructorStack()
	arr = []interface{}{}
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



