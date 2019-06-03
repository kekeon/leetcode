package stack

import "fmt"

type Stack struct {
	stack []interface{}
}

/** Initialize your data structure here. */
func ConstructorStack() Stack {

	return Stack{
		stack: []interface{}{},
	}
}

/** Push element x onto stack. */
func (this *Stack) Push(x interface{}) {
	this.stack = append(this.stack, x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *Stack) Pop() interface{} {
	size := len(this.stack)
	res := this.stack[size-1]
	this.stack = append([]interface{}{}, this.stack[0:size-1]...)
	return res
}

/** Get the top element. */
func (this *Stack) Top() interface{} {
	size := len(this.stack)

	if size == 0 {
		return -1
	}

	return this.stack[size-1]
}

/** Returns whether the stack is empty. */
func (this *Stack) Empty() bool {

	if len(this.stack) == 0 {
		return true
	} else {
		return false
	}
}

/**
 * Your Stack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func testStack() {
	obj := ConstructorStack()
	obj.Push(9)
	param_2 := obj.Pop()
	param_3 := obj.Top()
	param_4 := obj.Empty()

	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}
