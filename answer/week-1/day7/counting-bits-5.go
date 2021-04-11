
package day7

//思路不用每次循环去计算数目 通过递推获取目标值的数目
func countBits(num int) []int {
	var count = make([]int, num+1)
	count[0] = 0
	for i := 1; i <= num; i++ {
		count[i] = count[i&(i-1)] + 1
	}
	return count
}
