package main

import "fmt"

type Stack []int

func main() {
	my := Constructor()
	my.Push(1)
	my.Push(2)
	fmt.Println(my.Pop())
}

type MyQueue struct {
	in  *Stack
	out *Stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		&Stack{},
		&Stack{},
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	*this.in = append(*(this.in), x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(*(this.out)) == 0 {
		n := len(*(this.in)) //不能在条件内直接使用len因为len是变化的
		for i := 0; i < n; i++ {
			*(this.out) = append(*(this.out), (*(this.in)).Pop())
		}
		return this.out.Pop()
	} else {
		return this.out.Pop()
	}
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(*(this.out)) == 0 {
		n := len(*(this.in)) //不能在条件内直接使用len因为len是变化的
		for i := 0; i < n; i++ {
			*(this.out) = append(*(this.out), (*(this.in)).Pop())
		}
		return (*(this.out))[len(*(this.out))-1]
	} else {
		return (*(this.out))[len(*(this.out))-1]
	}
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(*(this.in)) == 0 && len(*(this.out)) == 0
}

func (s *Stack) Pop() int {
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}
