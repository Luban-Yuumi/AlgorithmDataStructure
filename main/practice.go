package main

import "math"

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	dp := make([][]int, n+1)
	for v := range dp {
		dp[v] = make([]int, m+1)
	}
	for i := 0; i <= n; i++ {
		dp[i][0] = i
	}
	for i := 0; i <= m; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word2[j-1] == word1[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = int(math.Min(float64(dp[i][j-1]), math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1])))) + 1
			}
		}
	}
	return dp[n][m]
}
