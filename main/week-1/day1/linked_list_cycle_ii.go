package main

import (
	"fmt"
)

func main() {
	a, b, c, d := new(ListNode), new(ListNode), new(ListNode), new(ListNode)
	a.Next = b
	b.Next = c
	c.Next = d
	d.Next = a

	fmt.Println(detectCycle(a))
}

//  链接：https://leetcode-cn.com/problems/linked-list-cycle-ii/solution/huan-xing-lian-biao-ii-by-leetcode-solution/
func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}


