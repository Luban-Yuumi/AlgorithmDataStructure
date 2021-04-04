package day8
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