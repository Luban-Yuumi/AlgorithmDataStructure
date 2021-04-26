package day9

import "strconv"

func numDecodings(s string) int {
	n := len(s)

	dp := make([]int, n+1)
	dp[0] = 1

	for i := 1; i <= n; i++ {
		if i > 1   {
			prev ,_ := strconv.Atoi(string(s[i-2]))
			double,_ := strconv.Atoi(s[i-2:i])
			if prev != 0 && double <=26 {
				if string(s[i-1]) != "0" {
					dp[i] = dp[i-1]+dp[i-2]
				}else {
					dp[i] = dp[i-2]
				}
				continue
			}
		}
		if string(s[i-1]) != "0" {
			dp[i] = dp[i-1]
		}
	}
	return dp[n]
}
