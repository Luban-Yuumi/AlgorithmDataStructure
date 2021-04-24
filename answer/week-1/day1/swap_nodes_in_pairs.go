package main

import "fmt"

//给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
//
//示例:
//
//给定 1->2->3->4, 你应该返回 2->1->4->3.
//说明:
//
//你的算法只能使用常数的额外空间。
//你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

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
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil && head.Next != nil{
		head.Next,head.Next.Next,head,prev.Next,prev = head.Next.Next,head,head.Next.Next,head.Next,head
	}
	return hair.Next
}
