package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

//方法一：暴力枚举 时间复杂度 （o(n3)）
//方法二：先枚举x,y 再判断-x-y是否在map中 时间复杂度（o(n2)）
func threeSum(nums []int) [][]int {
	res := [][]int{}
	var sumMap = make(map[int]int)
	var newMap = make(map[[3]int]int)
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums) //方便去重
	for i := 0; i < len(nums); i++ {
		sumMap = map[int]int{}
		for j := i + 1; j < len(nums); j++ {
			if _, ok := sumMap[nums[j]]; !ok {
				sumMap[-nums[i]-nums[j]] = 1
			} else {
				if _, ok := newMap[[3]int{nums[j], -nums[i] - nums[j], nums[i]}]; !ok { //去重
					res = append(res, []int{nums[j], -nums[i] - nums[j], nums[i]})
				}
				newMap[[3]int{nums[j], -nums[i] - nums[j], nums[i]}] = 1
			}
		}
	}

	return res
}
