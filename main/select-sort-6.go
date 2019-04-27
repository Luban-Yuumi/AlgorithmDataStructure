package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	sort(nums)
	fmt.Println(nums)
}
//选择排序(Selection-sort)是一种简单直观的排序算法。
// 它的工作原理：首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
// 然后，再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
// 以此类推，直到所有元素均排序完毕。时间复杂度为o(n2)
func sort(nums []int)  {
	if len(nums)<=1 {
		return
	}
	for i:=0;i<len(nums)-1 ;i++  {
		max := i
		for j:=i+1;j<len(nums) ;j++  {
			if nums[j]>nums[max] {
				max = j
			}
		}
		if max != i {
			nums[max],nums[i] = nums[i],nums[max]
		}
	}
}