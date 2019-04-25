package main

import (
	"math"
)

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	currentMax, currentMin, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		currentMax, currentMin = int(math.Max(math.Max(float64(currentMax*nums[i]), float64(currentMin*nums[i])), float64(nums[i]))), int(math.Min(math.Min(float64(currentMax*nums[i]), float64(currentMin*nums[i])), float64(nums[i])))
		if currentMax > res {
			res = currentMax
		}
	}
	return res
}
