package tree

import "fmt"

// 1. 中序遍历， 在生成树
func increasingBST(root *TreeNode1) *TreeNode1 {
	a := inOrder(root)
	t := TreeNode1{
		Val:a[0],
		Left:nil,
		Right:nil,
	}
	p := &t

	c := a[1:]
	for _, v := range c {
		b := TreeNode1{
			Val:v,
			Left:nil,
			Right:nil,
		}
		p.Right = &b
		p = &b
	}
	return &t
}

func getBranch(node *TreeNode1) []*TreeNode1 {
	list := []*TreeNode1{}
	for node != nil {
		list = append(list, node)
		node = node.Left
	}
	fmt.Println()
	return list
}

func inOrder(node *TreeNode1) []int {

	list := []*TreeNode1{}
	valList := []int{}
	for  {
		list = append(list, getBranch(node)...)
		l := len(list)
		if l == 0 {
			break
		}
		node  = list[l-1]
		list = list[0:l-1]
		valList = append(valList, node.Val)
		node = node.Right
	}

	return valList
}


// 2. 中序遍历， 将树重新拼接

func increasingBST2(root *TreeNode1) *TreeNode1 {
	a := inOrderTree(root)

	r := *&a[0]

	if r.Right == nil {
		r.Right = &TreeNode1{}
	}

	d := r

	for _, v := range a[1:] {
		d.Right = v
		if *&v.Right == nil {
			*&v.Right = &TreeNode1{}
		}
		d = v

	}
	return r
}

func getBranchNode(node *TreeNode1) []*TreeNode1 {
	list := []*TreeNode1{}
	for node != nil {
		list = append(list, node)
		node = node.Left
	}
	return list
}

func inOrderTree(node *TreeNode1) []*TreeNode1 {

	list := []*TreeNode1{}
	valList := []*TreeNode1{}
	for  {
		list = append(list, getBranchNode(node)...)
		l := len(list)
		if l == 0 {
			break
		}
		node = list[l-1]
		list = list[0:l-1]
		valList = append(valList, node)
		node = node.Right
	}

	return valList
}