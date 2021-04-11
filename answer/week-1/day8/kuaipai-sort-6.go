package day8

import "fmt"

//https://www.cnblogs.com/nima/p/12724868.html

func quickSort(nums []int, i, j int) {
	begin := partition(i, j, nums)
	if begin < 0 {
		return
	}

	quickSort(nums, i, begin-1)
	quickSort(nums, begin+1, j)

}

func partition(i, j int, nums []int) int {
	begin := i + 1
	end := j
	if begin > end {
		return -1
	}
	for begin < end {
		if nums[begin] > nums[i] {
			nums[begin], nums[end] = nums[end], nums[begin]
			end--
		} else {
			begin++
		}
	}

	if nums[begin] > nums[i] {
		begin = begin - 1
	}

	nums[i], nums[begin] = nums[begin], nums[i]
	return begin
}

func main() {
	list := []int{5}
	quickSort(list, 0, len(list)-1)
	fmt.Println(list)

	list1 := []int{5, 9}
	quickSort(list1, 0, len(list1)-1)
	fmt.Println(list1)

	list2 := []int{5, 9, 1}
	quickSort(list2, 0, len(list2)-1)
	fmt.Println(list2)

	list3 := []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3}
	quickSort(list3, 0, len(list3)-1)
	fmt.Println(list3)
}