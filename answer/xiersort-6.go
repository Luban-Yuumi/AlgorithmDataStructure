package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	sort(nums)
	fmt.Println(nums)
}
//先将整个待排序的记录序列分割成为若干子序列分别进行直接插入排序，具体算法描述：
//选择一个增量序列t1，t2，…，tk，其中ti>tj，tk=1；
//按增量序列个数k，对序列进行k 趟排序；
//每趟排序，根据对应的增量ti，将待排序列分割成若干长度为m 的子序列，分别对各子表进行直接插入排序。
// 仅增量因子为1 时，整个序列作为一个表来处理，表长度即为整个序列的长度。算法复杂度o(nlog2n)
func sort(nums []int)  {
	if len(nums)<=1 {
		return
	}
	incre := len(nums)
	for incre>=1  {
		incre/=2
		for i:=0;i<incre ;i++  {
			for j:=i+incre;j<len(nums) ;j+=incre  {
				for k:=j;k>i;k-=incre  {
					if nums[k] >nums[k-incre] {
						nums[k],nums[k-incre] = nums[k-incre],nums[k]
					}
				}
			}
		}
	}
}