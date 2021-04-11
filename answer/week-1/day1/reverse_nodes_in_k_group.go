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
	if head == nil || k <= 1 {
		return head
	}
	var slice []*ListNode
	for head != nil {
		slice, head = append(slice, head), head.Next
	}

	num := len(slice)
	if num < k {
		return slice[0]
	}

	result := slice[k-1]
	start := 0
	for num >= k {
		if num >= 2*k {
			slice[start].Next = slice[start+2*k-1]
		} else if num == k {
			slice[start].Next = nil //注意
		} else {
			slice[start].Next = slice[start+k]
		}
		for i := start + k - 1; i > start; i-- {
			slice[i].Next = slice[i-1]
		}
		start = start + k
		num = num - k
	}
	return result
}
