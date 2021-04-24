package main

import "fmt"

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
	result := reverseKGroup(a, 2)
	fmt.Printf("%d->%d->%d->%d", result.Val, result.Next.Val, result.Next.Next.Val, result.Next.Next.Next.Val)
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil{
		tail := prev
		for i:= 0;i<k;i++{
			tail = tail.Next
			if tail == nil{
				return hair.Next
			}
		}
		head,tail = myReverse(head,tail)
		prev.Next = head
		prev = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse (head *ListNode,tail *ListNode)(*ListNode,*ListNode){
	prev := tail.Next
	now := head
	for prev != tail {
		now,now.Next,prev = now.Next,prev,now
	}
	return tail,head
}