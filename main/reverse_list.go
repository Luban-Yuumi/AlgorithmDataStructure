package main

import "fmt"

//单链表反转
type ListNode struct {
	Next  *ListNode
	Value int
}
var prev  *ListNode = nil

func main() {
	var l1,l2,l3  = new(ListNode), new(ListNode), new(ListNode)
	l1.Value = 1
	l1.Next = l2
	l2.Value = 2
	l2.Next = l3
	l3.Value = 3
	fmt.Printf("the now list : %v -> %v -> %v",l1.Value,l2.Value,l3.Value)
	head := reverseRecursive(l1)
	fmt.Printf("the now list : %v -> %v -> %v",head.Value,head.Next.Value,head.Next.Next.Value)
}

//递归
func reverseRecursive(head *ListNode) *ListNode {
	var cur  = head
	if cur == nil {

	}else {
		var next  = head.Next
		cur.Next = prev
		prev = cur
		reverseRecursive(next)
	}
	return prev
}

//迭代
func reverseIteration(head *ListNode)*ListNode {
	var cur *ListNode
	var prev *ListNode
	cur, prev = head,nil
	for cur != nil {
		cur.Next,cur,prev = prev,cur.Next,cur
	}
	return prev
}
