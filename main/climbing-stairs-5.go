package main

//爬楼梯问题  = fibonacci
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	last_one := 2
	last_two := 1
	for i := 2; i < n; i++ {
		last_one, last_two = last_one+last_two, last_one
	}
	return last_one
}
