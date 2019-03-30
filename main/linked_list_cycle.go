package main

type ListNode struct {
	Next  *ListNode
	Value int
}

func main() {
	
}
//思路一 限时0.5s 跑完了 就是没环 没跑完就是有环 该方法太二

//思路二 每次循环通过一个map对其进行存储 并且每一次跑都检查map有没有该值 有则return true 没有则存入 复杂度o(n)

//思路三 龟兔赛跑
func hasCycle(head *ListNode) bool {
	var fast,slow  = head,head
	for fast!= nil && slow!=nil && fast.Next != nil  {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}
	return false
}