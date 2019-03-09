package stack

import "fmt"

type MyStack struct {
	stack []interface{}
}


/** Initialize your data structure here. */
func constructor() MyStack {

	return MyStack {
		stack: []interface{}{},
	}
}


/** Push element x onto stack. */
func (this *MyStack) Push(x int)  {
	this.stack = append(this.stack, x)
}


/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() interface{} {
	size := len(this.stack)
	res := this.stack[size-1]
	this.stack = append([]interface{}{}, this.stack[0:size-1]...)
	return res
}


/** Get the top element. */
func (this *MyStack) Top() interface{} {
	size := len(this.stack)

	if size == 0 {
		return -1
	}

	return this.stack[size-1]
}


/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {

	if len(this.stack) == 0 {
		return true
	}else {
		return false
	}
}


/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */

func main() {
	 obj := constructor()
	 obj.Push(9)
	 param_2 := obj.Pop()
	 param_3 := obj.Top()
	 param_4 := obj.Empty()

	 fmt.Println(param_2)
	 fmt.Println(param_3)
	 fmt.Println(param_4)
}
