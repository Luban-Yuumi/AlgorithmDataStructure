package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	charrSort(nums)
	fmt.Println(nums)
}
//从第一个元素开始，该元素可以认为已经被排序；
//取出下一个元素，在已经排序的元素序列中从后向前扫描；
//如果该元素（已排序）大于新元素，将该元素移到下一位置；
//重复步骤3，直到找到已排序的元素小于或者等于新元素的位置；
//将新元素插入到该位置后；最佳o(n) 最坏o(n2)
func charrSort(nums []int)  {
	if len(nums)<=1 {
		return
	}
	for i:=0;i<len(nums);i++  {
		for j:=i;j>0 ;j--  {
			if nums[j] > nums[j-1] {
				nums[j],nums[j-1] = nums[j-1],nums[j]
			}else {
				break
			}
		}
	}
}