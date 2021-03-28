package main

import "fmt"

func main() {
	nums := []int{1, 3, 1, 2, 0, 5}
	result := maxSlidingWindow(nums, 3)
	fmt.Println(result)
}

//方法1 使用最大堆 堆大小为k 每次进来一个元素 就删除第一个进来的堆元素 然后维护堆的排序 总共进行n次算法复杂的nlogk
//方法二 双端队列 的形式具体见实现 算法复杂度n
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	var res, windows []int
	for i, v := range nums {
		if i >= k && windows[0] <= i-k {
			if len(windows) > 1 {
				windows = windows[1:]
			} else {
				windows = nil
			}
		}
		for len(windows) > 0 && v > nums[windows[len(windows)-1]] {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, i)
		if i >= k-1 {
			res = append(res, nums[windows[0]])
		}

	}
	return res
}

