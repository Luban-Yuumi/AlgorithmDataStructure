package day2

type stack []int

func (s *stack) Pop() (v int) {
	n := len(*s)
	if n == 0 {
		return
	}
	v = (*s)[n-1]
	*s = (*s)[:n-1]
	return
}

func (s *stack) Push(v int) {
	*s = append(*s, v)
}

type MyQueue struct {
	in  stack
	out stack
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.in.Push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.out) != 0 {
		return this.out.Pop()
	}
	for len(this.in) > 0 {
		this.out.Push(this.in.Pop())
	}
	return this.out.Pop()
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.out) != 0 {
		return this.out[len(this.out)-1]
	}
	for len(this.in) > 0 {
		this.out.Push(this.in.Pop())
	}
	return this.out[len(this.out)-1]
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
