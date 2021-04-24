package day4

//方法1直接枚举所有情况 每天都有买进买出两种情况 时间复杂度 （2n）
//方法二贪心算法 o(n)
//方法三动态规划 o(n)
func maxProfitEasy(prices []int) int {
	var money int
	for i, _ := range prices {
		if i > 0 && prices[i] > prices[i-1] {
			money += prices[i] - prices[i-1]
		}
	}
	return money
}
