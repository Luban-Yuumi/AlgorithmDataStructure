package main

// 单链表反转
var prev *ListNode = nil

// 递归
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

// 迭代
func reverseIteration(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		head.Next, head, prev = prev, head.Next, head
	}
	return prev
}
