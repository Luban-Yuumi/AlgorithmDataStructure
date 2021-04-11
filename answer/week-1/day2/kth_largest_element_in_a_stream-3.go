package day2

import (
	"container/heap"
)

type smallHeap []int

func (s smallHeap) Len() int {
	return len(s)
}

func (s smallHeap) Less(a, b int) bool {
	return s[a] < s[b]
}

func (s smallHeap) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s *smallHeap) Push(n interface{}) {
	v, _ := n.(int)
	*s = append(*s, v)
}

func (s *smallHeap) Pop() interface{} {
	n := len(*s)
	if n == 0 {
		return 0
	}
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}

type KthLargest struct {
	s   smallHeap
	num int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		num: k,
		s:   make(smallHeap, 0, len(nums)),
	}
	for _, v := range nums {
		kth.Add(v)
	}
	return kth
}

func (this *KthLargest) Add(val int) int {
	if len(this.s) < this.num {
		heap.Push(&this.s, val)
	} else if val > this.s[0] {
		heap.Pop(&this.s)
		heap.Push(&this.s, val)
	}
	return this.s[0]
}