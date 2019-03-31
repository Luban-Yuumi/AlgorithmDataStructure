package main

import "container/list"

type MyStack struct {
	l *list.List
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{list.New()}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x interface{}) {
	this.l.PushBack(x)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() interface{} {
	return this.l.Remove(this.l.Back())
}

/** Get the top element. */
func (this *MyStack) Top() interface{} {
	return this.l.Back().Value
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return this.l.Len() == 0
}
