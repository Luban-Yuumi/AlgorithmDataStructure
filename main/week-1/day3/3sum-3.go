package day3

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
	var res [][]int
	var uniqueMap = make(map[[3]int]struct{})
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {
		tmpMap := make(map[int]struct{})
		for j := i + 1; j < len(nums); j++ {
			if _, ok := tmpMap[-nums[i]-nums[j]]; !ok {
				tmpMap[nums[j]] = struct{}{}
				continue
			} else if _, ok = uniqueMap[[3]int{nums[i], -nums[i] - nums[j], nums[j]}]; ok {
				continue
			}
			uniqueMap[[3]int{nums[i], -nums[i] - nums[j], nums[j]}] = struct{}{}
			res = append(res, []int{nums[i], -nums[i] - nums[j], nums[j]})
		}
	}
	return res
}
