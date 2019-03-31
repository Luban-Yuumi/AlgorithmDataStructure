package main

import "fmt"

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
	var stack Stack
	hash_map := map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
	}
	for _, v := range s {
		if _, ok := hash_map[string(v)]; !ok {
			stack = append(stack, string(v))
		} else if len(stack) == 0 || stack.Pop() != hash_map[string(v)] {
			return false
		}
	}
	return len(stack) == 0
}
