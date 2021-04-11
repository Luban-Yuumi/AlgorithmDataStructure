package day2

type queues []int

func (s *queues) Pop() (v int) {
	n := len(*s)
	if n == 0 {
		return
	}
	v = (*s)[0]
	if n > 1 {
		*s = (*s)[1:]
	} else {
		*s = nil
	}
	return
}

func (s *queues) Push(v int) {
	*s = append(*s, v)
}

type MyStack struct {
	q1 queues
	q2 queues
}

/** Initialize your data structure here. */
func ConstructorStack() MyStack {
	return MyStack{}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.q2.Push(x)
	for len(this.q1) > 0 {
		this.q2.Push(this.q1.Pop())
	}
	this.q1, this.q2 = this.q2, this.q1
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	return this.q1.Pop()
}

/** Get the top element. */
func (this *MyStack) Top() int {
	if len(this.q1) == 0 {
		return 0
	}
	return this.q1[0]
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	return len(this.q1) == 0
}
