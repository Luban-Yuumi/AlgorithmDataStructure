package main

import (
	"fmt"
)

type ListNode struct {
	Next  *ListNode
	Value int
}

func main() {
	a, b, c, d := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = a

	fmt.Println(detectCycle(a))
}

func detectCycle(head *ListNode) *ListNode {
	a := map[*ListNode]int{}
	i := 0
	for head != nil && head.Next != nil {
		if _, ok := a[head]; ok {
			return head
		} else {
			a[head] = i
		}
		i++
		head = head.Next
	}
	return nil
}
