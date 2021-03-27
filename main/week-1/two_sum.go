package main

//方法一 暴力枚举 这种方法的算法复杂度为o(n2)
//方法二 如下 复杂度为o()/n
func twoSum(nums []int, target int) []int {
	sumMap := make(map[int]int)
	for k, v := range nums {
		if index, ok := sumMap[target-v]; ok {
			return []int{k, index}
		}
		sumMap[v] = k
	}
	return nil
}
