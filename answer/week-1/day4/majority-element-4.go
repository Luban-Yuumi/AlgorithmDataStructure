package day4

//方法一暴力枚举 两层循环 一层循环遍历数组 一层循环统计x的数目 时间复杂度O(n2)
//方法二通过sort排序后 再通过一层循环进行x数目统计 时间复杂度 O(nlogn)
//方法三通过map进行每个数组的数目存储 最后的出最大的 时间复杂度O(n)
//方法四通过分治的思维，每次将数组一分为二 然后分别求众数if left == right return left or right else return max(left,right)

func majorityElement(nums []int) int {
	var countMap = make(map[int]int)
	for _, v := range nums {
		countMap[v] += 1
	}
	var a int
	var key int
	for i, v := range countMap {
		if v > a {
			a = v
			key = i
		}
	}
	return key
}
