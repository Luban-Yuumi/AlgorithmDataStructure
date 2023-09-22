package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k < 2 {
		return head
	}

	var prev, result = head, head

	for head != nil {
		tail := head
		for i := 0; i < k-1; i++ {
			if tail == nil {
				return result
			}

			tail = tail.Next
		}

		if prev == nil {
			result = tail
		}

		if prev != nil {
			prev.Next = tail
		}

		prev = head
		reverse(head, tail)
		head = head.Next
	}

	return result

}

func reverse(head *ListNode, tail *ListNode) {
	prev := tail.Next
	cur := head

	for prev != tail {
		cur.Next, cur, prev = prev, cur.Next, cur
	}
}
