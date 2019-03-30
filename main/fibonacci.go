package main

import "fmt"

func main() {
	//recursive
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciRecursive(i))
	}
	//Iteration
	for i := 0; i < 10; i++ {
		fmt.Println(fibonacciRecursive(i))
	}
}

func fibonacciRecursive(i int) int {
	if i <= 1 {
		return i
	} else {
		return fibonacciRecursive(i-1) + fibonacciRecursive(i-2)
	}
}

func fibonacciIteration(i int) int {
	if i <= 0 {
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
