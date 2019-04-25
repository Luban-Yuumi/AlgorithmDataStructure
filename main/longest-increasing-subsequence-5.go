package main

func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	for i := 1; i < n; i++ {
		var max = 0
		for j := 0; j < i; j++ {
			if dp[j] > max && nums[j] < nums[i] {
				max = dp[j]
			}
		}
		dp[i] = max + 1
	}
	var max = 0
	for _, v := range dp {
		if v > max {
			max = v
		}
	}
	return max
}
