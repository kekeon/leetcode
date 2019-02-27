package main

import "fmt"

type MyQueue struct {
	queue []interface{}
}


/** Initialize your data structure here. */
func Constructor() MyQueue {

	return MyQueue{
		queue:[]interface{}{},
	}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x interface{})  {
	this.queue = append(this.queue,x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() interface{} {

	size := len(this.queue)

	if size == 0 {
		return  nil
	}else {
		res := this.queue[0]
		this.queue = append([]interface{}{},this.queue[1:]...)
		return res
	}
}


/** Get the front element. */
func (this *MyQueue) Peek() interface{} {
	size := len(this.queue)
	if size == 0 {
		return  nil
	}else {
		return this.queue[0]
	}
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {

	if len(this.queue) == 0 {
		return  true
	}else {
		return false
	}
}


/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */

func main() {
	obj := Constructor()
	obj.Push("12")
	param_2 := obj.Pop()
	param_3 := obj.Peek()
	param_4 := obj.Empty()

	fmt.Println(param_2)
	fmt.Println(param_3)
	fmt.Println(param_4)
}
