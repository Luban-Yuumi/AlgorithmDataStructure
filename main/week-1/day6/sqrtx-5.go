package day6

import "fmt"

func main() {
	fmt.Println(mySqrtWithAccuracy(3, 0.001))
}

//方法1 二分法
func mySqrt1(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	left, right := 1, x
	var res int
	for right-left >= 0 {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - 1
		} else {
			left = mid + 1
			res = mid
		}
	}
	return res
}

//方法2牛顿迭代法
func mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}
	res := x
	for res*res > x {
		res = (res + x/res) / 2
	}
	return res
}

//带精度的

func mySqrtWithAccuracy(x float64, accuracy float64) float64 {
	if x == 0 {
		return x
	}
	var left float64 = 0
	var right = x
	for right-left >= accuracy {
		mid := (left + right) / 2
		if mid*mid == x {
			return mid
		} else if mid*mid > x {
			right = mid - accuracy
		} else {
			left = mid + accuracy
		}
	}
	return right
}
