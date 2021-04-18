package day8

func sortArray(nums []int) []int {
	//quickSort(nums, 0, len(nums)-1)
	maoPaoSort(nums)
	return nums
}

func quickSort(nums []int, begin, end int) {
	if begin >= end {
		return
	}
	i := begin + 1
	j := end
	for j > i {
		if nums[i] > nums[begin] {
			nums[i], nums[j] = nums[j], nums[i]
			j--
		} else {
			i++
		}
	}
	mid := 0
	if nums[i] > nums[begin] {
		mid = i - 1
	} else {
		mid = i
	}
	nums[begin], nums[mid] = nums[mid], nums[begin]
	quickSort(nums, begin, mid-1)
	quickSort(nums, mid+1, end)
}

func maoPaoSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		change := false
		for j := 0; j < len(nums)-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				change = true
			}
		}

		if !change {
			break
		}
	}
}
