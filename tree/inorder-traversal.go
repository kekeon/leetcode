package tree

import "leetcode/stack"

// 中序遍历

func goAlongLeftBranch(x *TreeNode, s *stack.Stack) {
	for x != nil {
		s.Push(x)
		x = x.Left
	}
}


func InorderTraversal(x *TreeNode) []interface{} {

	arr := []interface{}{}  //将所有中序访问的值存储

	stackNode := stack.ConstructorStack()

	for {
		goAlongLeftBranch(x,&stackNode)  //从当前节点出发，将所有的做左子树放入栈中

		if stackNode.Empty() {  // 直到栈节点处理完毕
			break
		}

		x = stackNode.Pop().(*TreeNode)  // 每当一个左子树，处理完毕后，将访问权，交给下一层左子树
		arr = append(arr,x.Val)
		x = x.Right  // 转给下一个左子树
	}

	return arr

}