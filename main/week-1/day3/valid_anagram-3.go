package day3

import "reflect"

//方法一：通过排序函数 然后再做对比 复杂度Nlogn
//方法二如下
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	var (
		mapS = make(map[rune]int)
		mapT = make(map[rune]int)
	)
	for _, v := range s {
		mapS[v] += 1
	}
	for _, v := range t {
		mapT[v] += 1
	}

	return reflect.DeepEqual(mapS, mapT)
}

