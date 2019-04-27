package main

import "fmt"

func main() {
	nums := []int{1,2,3,4,5}
	kuaisu(nums)
	fmt.Println(nums)
}
//快速排序使用分治法来把一个串（list）分为两个子串（sub-lists）。具体算法描述如下：
//从数列中挑出一个元素，称为 “基准”（pivot）；
//重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作；
//递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序。最佳情况：T(n) = O(nlogn)   最差情况：T(n) = O(n2)   平均情况：T(n) = O(nlogn)　
func kuaisu(buf []int) {
	kuai(buf, 0, len(buf)-1)
}

func kuai(a []int, l, r int) {
	if l>=r {
		return
	}
	i,j,k := l,r,a[l]
	for i<j {
		for i<j&&a[j]<=k  {
			j--
		}
		if i<j {
			a[i] = a[j]
			i++
		}
		for i<j&&a[i]>=k  {
			i++
		}
		if i<j  {
			a[j] = a[i]
			j--
		}
	}
	a[i] = k
	kuai(a,l,i-1)
	kuai(a,i+1,r)
}