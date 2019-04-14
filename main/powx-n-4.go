package main

import "fmt"

//方法一暴力循环乘法,时间复杂度o（n）
//方法二 分治思想 时间复杂度o(logn)
//递归
func myPowResursive(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 1 / myPowResursive(x, -n)
	}
	if n%2 > 0 {
		return x * myPowResursive(x, n-1)
	}
	return myPowResursive(x*x, n/2)
}

//迭代
func myPowIteration(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		x = 1 / x
		n = -n
	}
	pow := 1.0
	for n > 0 {
		if n&1 > 0 { //奇数多乘个s
			pow *= x
		}
		x *= x
		n >>= 1
	}
	return pow
}

func main() {
	fmt.Println(myPowIteration(2.0, 5))
}
