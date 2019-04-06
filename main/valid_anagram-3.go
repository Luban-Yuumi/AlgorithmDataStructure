package main

import "reflect"

//方法一：通过排序函数 然后再做对比 复杂度Nlogn
func isAnagram(s string, t string) bool {
	map1 := make(map[string]int)
	map2 := make(map[string]int)
	for _, v := range s {
		map1[string(v)] += 1
	}
	for _, v := range t {
		map2[string(v)] += 1
	}
	return reflect.DeepEqual(map1, map2)
}
