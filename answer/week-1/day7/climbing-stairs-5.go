package day7

//爬楼梯问题  = fibonacci
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	prevTwo := 1
	prevOne := 2
	for i := 2; i < n; i++ {
		prevOne, prevTwo = prevOne+prevTwo, prevOne
	}
	return prevOne
}

