package main

//方法1直接枚举所有情况 每天都有买进买出两种情况 时间复杂度 （2n）
//方法二贪心算法 o(n)
//方法三动态规划 o(n)
func maxProfit(prices []int) int {
	var money int
	for i, v := range prices {
		if i != 0 && v > prices[i-1] {
			money += v - prices[i-1]
		}
	}
	return money
}
