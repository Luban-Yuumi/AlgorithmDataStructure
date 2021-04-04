package day8

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	sort(nums)
	fmt.Println(nums)
}
//比较相邻的元素。如果第一个比第二个大，就交换它们两个；
//对每一对相邻元素作同样的工作，从开始第一对到结尾的最后一对，这样在最后的元素应该会是最大的数；
//针对所有的元素重复以上的步骤，除了最后一个；
//重复步骤1~3，直到排序完成。 时间复杂度最好为o(n)最坏为o(n2)
func sort(nums []int)  {
	if len(nums)<=1 {
		return
	}
	for i:=0;i<len(nums)-1 ;i++  {
		flag := false
		for j:=0;j<len(nums)-i-1 ;j++  {
			if  nums[j]<nums[j+1]{
				nums[j],nums[j+1] = nums[j+1],nums[j]
			}
			flag = true
		}
		if !flag {
			break
		}
	}
}