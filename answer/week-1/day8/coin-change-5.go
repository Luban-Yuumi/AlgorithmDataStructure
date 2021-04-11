package day8

import "math"

func coinChange(coins []int, amount int) int {
	if amount == 0 || len(coins) == 0 {
		return -1
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
		for _, v := range coins {
			if i >= v {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-v]+1)))
			}
		}
	}
	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}
