package main

import (
	"math"
	"strconv"
)

//方法一递归搜索 时间复杂度 o（2n） d[i,j] = min(d[i+1,j],d[i+1,j+1])
//方法二贪心算法 这种方法是不适合的会出错
//方法三动态规划 从最后一层往最上面一层递推 求最优解
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	if len(triangle) == 1 {
		return triangle[0][0]
	}
	var dp = make(map[string]int)
	for i := len(triangle) - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == len(triangle)-1 {
				dp[strconv.Itoa(i)+","+strconv.Itoa(j)] = triangle[i][j]
			} else {
				dp[strconv.Itoa(i)+","+strconv.Itoa(j)] = int(math.Min(float64(dp[strconv.Itoa(i+1)+","+strconv.Itoa(j)]), float64(dp[strconv.Itoa(i+1)+","+strconv.Itoa(j+1)]))) + triangle[i][j]
			}
		}
	}
	return dp["0,0"]
}
