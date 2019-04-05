package main

import (
	"fmt"
)

type Stack []string

func (s *Stack) Pop() string {
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}

func main() {
	s := "()"
	fmt.Println(isValid(s))
}

func isValid(s string) bool {
	var stack  = Stack{}
	var stringMap = map[string]string{
		"}":"{",
		"]":"[",
		")":"(",
	}
	for _,s := range s{
		if _,ok := stringMap[string(s)];!ok {
			stack = append(stack,s)
		}else {
			if len(stack) == 0 || stringMap[string(s)]!= stack.Pop() {
				return  false
			}
		}
	}
	return len(stack) == 0
}
