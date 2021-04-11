package day7

//爬楼梯问题  = fibonacci
func climbStairs(n int) int {
	var dp = make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i == 1 {
			dp[0] = 1
			dp[1] = 1
		} else {
			dp[i] = dp[i-1] + dp[i-2]
		}
	}
	return dp[n]
}
