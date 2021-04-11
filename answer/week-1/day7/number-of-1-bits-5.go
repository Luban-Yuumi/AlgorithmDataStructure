package day7

import (
	"fmt"
)

// x & (x-1) 删除最后一位的1 x & -x 获得最后一位的1

//方法1 二进制每次与2取余数 有余数计数器加1然后数字往右移一位 由于int类型的最大长度已知 所以时间复杂度 为o(1)

func hammingWeight1(num uint32) int {
	count := 0
	for num != 0 {
		if num&1 == 1 {
			count++
		}
		num >>= 1
	}
	return count
}

//方法二 每次去除最低位1
func hammingWeight(num uint32) int {
	count := 0
	for num != 0 {
		num = num & (num - 1)
		count++
	}
	return count
}

func main() {
	fmt.Println(hammingWeight(-1))
}
