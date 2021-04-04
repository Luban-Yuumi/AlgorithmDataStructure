package day8

import (
	"math"
)

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	DP := [2][2]int{{nums[0], nums[0]}}
	res := DP[0][0]
	for i := 1; i < len(nums); i++ {
		now, prev := i%2, (i-1)%2
		DP[now][0] = int(math.Max(math.Max(float64(DP[prev][0]*nums[i]), float64(DP[prev][1]*nums[i])), float64(nums[i])))
		DP[now][1] = int(math.Min(math.Min(float64(DP[prev][0]*nums[i]), float64(DP[prev][1]*nums[i])), float64(nums[i])))
		res = int(math.Max(float64(DP[now][0]), float64(res)))
	}

	return res
}

