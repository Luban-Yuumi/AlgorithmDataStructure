package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	var a, b, c, d = new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a.Val = 1
	a.Next = b
	b.Val = 2
	b.Next = c
	c.Val = 3
	c.Next = d
	d.Val = 4
	fmt.Printf("%d->%d->%d->%d\n", a.Val, b.Val, c.Val, d.Val)
	result := swapPairs(a)
	fmt.Printf("%d->%d->%d->%d", result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val)
}

func swapPairs(head *ListNode) *ListNode {
	var pre = new(ListNode)
	if head == nil || head.Next == nil {
		return head
	}
	newHead := head.Next
	for head != nil && head.Next != nil {
		a := head
		b := head.Next
		pre.Next, a.Next, b.Next, pre = b, b.Next, a, a
		head = a.Next
	}

	return newHead
}
