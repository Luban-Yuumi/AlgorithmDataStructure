package day8

import "math"

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}

	dp, res := make([]int, n), 0

	for i := 0; i < n; i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[i]), float64(dp[j]+1)))
			}
		}
		res = int(math.Max(float64(dp[i]), float64(res)))
	}
	return res
}
