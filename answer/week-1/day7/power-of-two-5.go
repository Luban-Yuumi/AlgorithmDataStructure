package day7

//2的n次幂代表只有一个1 去除1后是等于0的
func isPowerOfTwo(n int) bool {
	return n&(n-1) == 0 && n != 0
}

