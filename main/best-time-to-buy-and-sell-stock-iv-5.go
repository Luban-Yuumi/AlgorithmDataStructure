package main

import (
	"fmt"
	"math"
)

//动态规划 涉及到三个状态 天数，是否持有股票，交易次数
func maxProfit(k int, prices []int) int {
	if len(prices) <= 1 || k <= 1 {
		return 0
	}
	days := len(prices)
	var dep = make([][][]int, days)
	for i := range dep {
		dep[i] = make([][]int, 2)
		for j := range dep[i] {
			dep[i][j] = make([]int, k+1)
		}
	}
	for i := 0; i <= k; i++ {
		dep[0][0][k] = 0
		if i < k {
			dep[0][1][i] = -prices[0]
		}
	}
	for i := 1; i < days; i++ {
		for j := 0; j <= k; j++ {
			if j == 0 {
				dep[i][0][0] = 0
				dep[i][1][0] = -prices[i]
			} else {
				dep[i][0][j] = int(math.Max(float64(dep[i-1][1][j-1]+prices[i]), float64(dep[i-1][0][j])))
				dep[i][1][j] = int(math.Max(float64(dep[i-1][0][j]-prices[i]), float64(dep[i-1][1][j])))
			}
		}
	}
	res := 0
	for _, v := range dep[days-1][0] {
		if v > res {
			res = v
		}
	}
	return res
}
func main() {
	k := 2
	price := []int{1, 2, 4, 5, 9, 8, 3}
	fmt.Println(maxProfit(k, price))
}
