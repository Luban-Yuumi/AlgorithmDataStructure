package day7

import (
	"math"
)

//方法一递归搜索 时间复杂度 o（2n） d[i,j] = min(d[i+1,j],d[i+1,j+1])
//方法二贪心算法 这种方法是不适合的会出错
//方法三动态规划 从最后一层往最上面一层递推 求最优解
func minimumTotal(triangle [][]int) int {
	if len(triangle) == 0 {
		return 0
	}
	n := len(triangle)
	var res = make([]int, len(triangle[n-1]))
	for i := n - 1; i >= 0; i-- {
		for j := 0; j < len(triangle[i]); j++ {
			if i == n-1 {
				res[j] = triangle[i][j]
			} else {
				res[j] = int(math.Min(float64(res[j]), float64(res[j+1]))) + triangle[i][j]
			}
		}
	}
	return res[0]
}
