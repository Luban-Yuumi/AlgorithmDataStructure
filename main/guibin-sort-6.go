package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	sort(nums)
	fmt.Println(nums)
}
//把长度为n的输入序列分成两个长度为n/2的子序列；
//对这两个子序列分别采用归并排序；
//将两个排序好的子序列合并成一个最终的排序序列。
func sort(nums []int)  {
	if len(nums)<=1 {
		return
	}

}