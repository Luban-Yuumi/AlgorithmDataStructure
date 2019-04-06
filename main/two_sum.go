
package main

import (
	"fmt"
	"sync"
)

var (
	nums   = []int{2, 7, 11, 15}
	noSolu = []int{-1, -1}
	target = 26
	wg     sync.WaitGroup
)

type Num struct {
	num, index int
}

type Nums []Num

func (slice Nums) Len() int {
	return len(slice)
}

func (slice Nums) Less(i, j int) bool {
	return slice[i].num < slice[j].num
}

func (slice Nums) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// 普通暴力 O(N^2)
func algo1() []int {
	size := len(nums)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return noSolu
}

// O(1) 算法（滑稽
func algo4() (ret []int){
	ret = noSolu
	size := len(nums)
	wg.Add((size - 1) * size / 2)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			go func(i, j int) {
				if nums[i] + nums[j] == target {
					ret = []int{i, j}
				}
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	return
}

func main() {
	fmt.Println(algo1())
	fmt.Println(algo2())
	fmt.Println(algo3())
	fmt.Println(algo4())
}