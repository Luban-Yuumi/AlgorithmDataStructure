package main

import "math"

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	var dp [2][2]int
	dp[0][0] = nums[0]
	dp[0][1] = nums[0]
	var res = dp[0][0]

	for i := 1; i < len(nums); i++ {
		x, y := i&1, (i-1)&1
		dp[x][0] = int(math.Max(math.Max(float64(dp[y][0]*nums[i]), float64(dp[y][1]*nums[i])), float64(nums[i])))
		dp[x][1] = int(math.Min(math.Min(float64(dp[y][0]*nums[i]), float64(dp[y][1]*nums[i])), float64(nums[i])))
		res = int(math.Max(float64(res), float64(dp[x][0])))
	}

	return res
}
