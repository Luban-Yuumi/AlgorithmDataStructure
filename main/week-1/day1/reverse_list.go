package main

import "fmt"

//单链表反转
var prev *ListNode = nil

func main() {
	var l1, l2, l3 = new(ListNode), new(ListNode), new(ListNode)
	l1.Val = 1
	l1.Next = l2
	l2.Val = 2
	l2.Next = l3
	l3.Val = 3
	fmt.Printf("the now list : %v -> %v -> %v", l1.Val, l2.Val, l3.Val)
	head := reverseRecursive(l1)
	fmt.Printf("the now list : %v -> %v -> %v", head.Val, head.Next.Val, head.Next.Next.Val)
}

//递归
func reverseRecursive(head *ListNode) *ListNode {
	var cur = head
	if cur == nil {
		return prev
	}
	var next = head.Next
	cur.Next = prev
	prev = cur
	reverseRecursive(next)
	return prev
}

//迭代
func reverseIteration(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}
