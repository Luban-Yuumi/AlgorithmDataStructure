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
	cur := head
	if k <= 1 || head == nil {
		return head
	}
	var slice []*ListNode
	i := 0
	for head != nil {
		slice = append(slice, head)
		head = head.Next
	}
	num := len(slice)
	start := 0
	if num < k {
		return cur //head 后面被覆盖了 修复
	} else {
		newHead := slice[k-1]
		for num >= k {
			if num-k >= k {
				slice[start].Next = slice[start+(2*k)-1] //如果后面要反转的话，首节点的next就会是反转之后的节点
			} else {
				if num == k {
					slice[start].Next = nil //修复index out of range
				} else {
					slice[start].Next = slice[start+k]
				}
			}

			for i := start + 1; i < k+start; i++ {
				slice[i].Next = slice[i-1]
			}
			num = num - k
			start = start + k
		}
		return newHead
	}

}
