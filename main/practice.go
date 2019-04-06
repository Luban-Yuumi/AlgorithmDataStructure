package main

type Stack []string

func (s *Stack) Pop() string {
	n := len(*s)
	v := (*s)[n-1]
	*s = (*s)[:n-1]
	return v
}

func isValid(s string) bool {
	var slice = Stack{}
	var kMap = map[string]string{"}": "{", "]": "[", ")": "("}
	for _, v := range s {
		if _, ok := kMap[string(v)]; !ok {
			slice = append(slice, string(v))
		} else if len(slice) == 0 || slice.Pop() != kMap[string(v)] {
			return false
		}
	}
	return len(slice) == 0
}
