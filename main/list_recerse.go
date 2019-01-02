package main

import "fmt"

type ListNode struct {
	value int
	next  *ListNode
}

func main() {
	var a, b, c, d ListNode
	a.value = 1
	a.next = &b
	b.value = 2
	b.next = &c
	c.value = 3
	c.next = &d
	d.value = 4
	fmt.Printf("%d->%d->%d->%d\n", a.value, b.value, c.value, d.value)
	result := reverseList(&a)
	fmt.Printf("%d->%d->%d->%d", result.value, result.next.value, result.next.next.value, result.next.next.next.value)
}

//递归解法
func reverseList(head *ListNode) *ListNode {
	if head.next == nil || head == nil {
		return head
	} else {
		a := reverseList(head.next)
		head.next.next = head
		return a
	}
}

//迭代解法
//func reverseList(head *ListNode) *ListNode {
//	if head.next == nil || head == nil {
//		return head
//	} else {
//		b := head.next
//		a := b.next
//		b.next = head
//		head.next = nil
//		for {
//			if a != nil {
//				c := a.next
//				a.next = b
//				b = a
//				a = c
//			} else {
//				return b
//			}
//		}
//	}
//}
