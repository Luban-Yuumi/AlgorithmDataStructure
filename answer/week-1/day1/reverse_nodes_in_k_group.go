package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	hair := &ListNode{Next: head}
	prev := hair
	for head != nil {
		tail := prev
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return hair.Next
			}
		}
		head, tail = myReverse(head, tail)
		prev.Next = head
		prev = tail
		head = tail.Next
	}
	return hair.Next
}

func myReverse(head *ListNode, tail *ListNode) (*ListNode, *ListNode) {
	prev := tail.Next
	now := head
	for prev != tail {
		now, now.Next, prev = now.Next, prev, now
	}
	return tail, head
}
