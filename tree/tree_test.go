package tree

import (
	"testing"
	"fmt"
)



// go test -v .\tree -test.run TestPreorderTraversal
// 先序遍历测试
func TestPreorderTraversal(t *testing.T){

	gNode := TreeNode{
		Val: "g",
		Left:	nil,
		Right: nil,
	}

	fNode := TreeNode{
		Val: "f",
		Left:	&gNode,
		Right: nil,
	}

	eNode := TreeNode{
		Val: "e",
		Left:	nil,
		Right: nil,
	}



	bNode := TreeNode{
		Val: "b",
		Left:	nil,
		Right: nil,
	}

	dNode := TreeNode{
		Val: "d",
		Left:	nil,
		Right: &eNode,
	}

	cNode := TreeNode{
		Val: "c",
		Left:	&dNode,
		Right: &fNode,
	}



	aNode := TreeNode{
		Val: "a",
		Left:	&bNode,
		Right: &cNode,
	}



	p := travPre(&aNode)
	q := travPre(&aNode)
	fmt.Println(p)
	fmt.Println(q)
}

// 两二叉树是否相同
func TestIsSameTree(t *testing.T){


	gNode := TreeNode{
		Val: "g",
		Left:	nil,
		Right: nil,
	}

	fNode := TreeNode{
		Val: "f",
		Left:	&gNode,
		Right: nil,
	}

	eNode := TreeNode{
		Val: "e",
		Left:	nil,
		Right: nil,
	}



	bNode := TreeNode{
		Val: "b",
		Left:	nil,
		Right: nil,
	}

	dNode := TreeNode{
		Val: "d",
		Left:	nil,
		Right: &eNode,
	}

	cNode := TreeNode{
		Val: "c",
		Left:	&dNode,
		Right: &fNode,
	}



	aNode := TreeNode{
		Val: "a",
		Left:	&bNode,
		Right: &cNode,
	}

	v := isSameTree(&aNode, &aNode)

	fmt.Println(v)
}

// 中序遍历测试
func TestInorderTraversal(t *testing.T) {

	gNode := TreeNode{
		Val: "g",
		Left:	nil,
		Right: nil,
	}

	eNode := TreeNode{
		Val: "e",
		Left:	nil,
		Right: nil,
	}

	cNode := TreeNode{
		Val: "c",
		Left:	nil,
		Right: nil,
	}

	dNode := TreeNode{
		Val: "d",
		Left:	&cNode,
		Right: &eNode,
	}

	fNode := TreeNode{
		Val: "f",
		Left:	&dNode,
		Right: &gNode,
	}

	aNode := TreeNode{
		Val: "a",
		Left:	nil,
		Right: nil,
	}

	bNode := TreeNode{
		Val: "b",
		Left:	&aNode,
		Right: &fNode,
	}

	v := InorderTraversal(&bNode)

	fmt.Println(v)
}

// 层次遍历测试
func TestLevelTraversal(t *testing.T) {
	gNode := TreeNode{
		Val: "g",
		Left:	nil,
		Right: nil,
	}

	eNode := TreeNode{
		Val: "e",
		Left:	nil,
		Right: &gNode,
	}

	cNode := TreeNode{
		Val: "c",
		Left:	nil,
		Right: nil,
	}


	fNode := TreeNode{
		Val: "f",
		Left:	nil,
		Right: nil,
	}


	dNode := TreeNode{
		Val: "d",
		Left:	&eNode,
		Right: &fNode,
	}

	bNode := TreeNode{
		Val: "b",
		Left:	&cNode,
		Right: &dNode,
	}

	aNode := TreeNode{
		Val: "a",
		Left:	&bNode,
		Right: nil,
	}


	v := LevelTraversal(&aNode)

	fmt.Println(v)


}

func TestLevelOrderBottom(t *testing.T){

	gNode := TreeNode1{
		Val: "g",
		Left:	nil,
		Right: nil,
	}

	eNode := TreeNode1{
		Val: "e",
		Left:	nil,
		Right: &gNode,
	}




	fNode := TreeNode1{
		Val: "f",
		Left:	nil,
		Right: nil,
	}


	dNode := TreeNode1{
		Val: 15,
		Left:	nil,
		Right: nil,
	}

	cNode := TreeNode1{
		Val: 20,
		Left:	&dNode,
		Right: &eNode,
	}

	bNode := TreeNode1{
		Val: 9,
		Left:	nil,
		Right: nil,
	}

	aNode := TreeNode1{
		Val: 3,
		Left:	&bNode,
		Right: &cNode,
	}

	v:= levelOrderBottom(aNode)
	fmt.Println(v)
}