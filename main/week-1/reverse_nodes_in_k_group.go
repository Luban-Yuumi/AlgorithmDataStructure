package main

import "fmt"

type ListNode struct {
	Next  *ListNode
	Value int
}

func main() {
	var a, b, c, d = new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a.Value = 1
	a.Next = b
	b.Value = 2
	b.Next = c
	c.Value = 3
	c.Next = d
	d.Value = 4
	fmt.Printf("%d->%d->%d->%d\n", a.Value, b.Value, c.Value, d.Value)
	result := reverseKGroup(a, 2)
	fmt.Printf("%d->%d->%d->%d", result.Value, result.Next.Value, result.Next.Next.Value, result.Next.Next.Next.Value)
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

