package main

import "math"

func coinChange(coins []int, amount int) int {
	if amount < 1 {
		return amount
	}
	//先将初始值全设置为amount+1方便后面求min
	dp := make([]int, amount+1)
	for j := range dp {
		if j > 0 {
			dp[j] = amount + 1
		}
	}
	for i := 1; i <= amount; i++ {
		for _, v := range coins {
			if i-v >= 0 {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-v])+1))
			}
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}
