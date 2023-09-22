package main

func fibonacciRecursive(i int) int {
	if i <= 1 {
		return i
	} else {
		return fibonacciRecursive(i-1) + fibonacciRecursive(i-2)
	}
}

func fibonacciIteration(i int) int {
	if i <= 1 {
		return i
	} else {
		a := 0
		b := 1
		for n := 2; n <= i; n++ {
			a, b = a+b, a
		}
		return a
	}
}
