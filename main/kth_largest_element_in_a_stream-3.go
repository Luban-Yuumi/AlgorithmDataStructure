package main

import "container/heap"

type smallHeap []int

func (s smallHeap) Len() int {
	return len(s)
}
func (s smallHeap) Swap(a, b int) {
	s[a], s[b] = s[b], s[a]
}

func (s smallHeap) Less(a, b int) bool {
	return s[a] < s[b]
}
func (s *smallHeap) Push(a interface{}) {
	*s = append(*s, a.(int))
}
func (s *smallHeap) Pop() interface{} {
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}

type KthLargest struct {
	s   *smallHeap
	num int
}

func Constructor(k int, nums []int) KthLargest {
	var kthLargest = KthLargest{
		s:   &smallHeap{},
		num: k,
	}
	for i := 0; i < len(nums); i++ {
		(&kthLargest).Add(nums[i])
	}
	return kthLargest
}

//每次堆插入的算法复杂度 为logn总共插入n次 花时n*logn
func (this *KthLargest) Add(val int) int {
	if len(*this.s) < this.num {
		heap.Push(this.s, val)
	} else if val > (*this.s)[0] {
		heap.Pop(this.s)
		heap.Push(this.s, val)
	}
	return (*this.s)[0]
}

//还有一种方法为每次将k个存入数组 每次都做一次快排 然后取出第一个 这种方法为n * klogk
