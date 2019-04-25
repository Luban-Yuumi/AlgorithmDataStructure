package main

import (
	"math"
)

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	dep := make([][]int, n+1)
	for v := range dep {
		dep[v] = make([]int, m+1)
	}
	for v := range dep {
		dep[v][0] = v
	}
	for v := range dep[0] {
		dep[0][v] = v
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if word1[i-1] == word2[j-1] {
				dep[i][j] = dep[i-1][j-1]
			} else {
				dep[i][j] = int(math.Min(float64(dep[i][j-1]), math.Min(float64(dep[i-1][j]), float64(dep[i-1][j-1])))) + 1
			}
		}
	}
	return dep[n][m]
}
