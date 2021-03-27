package main

type stack []string

func (s *stack) Pop() (v string) {
	n := len(*s)
	if n == 0 {
		return
	}
	v = (*s)[n-1]
	*s = (*s)[:n-1]
	return
}

func (s *stack) Push(v string) {
	*s = append(*s, v)
}

func isValid(s string) bool {
	var myStack stack
	parenMap := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	for _, v := range s {
		if _, ok := parenMap[string(v)]; ok {
			myStack.Push(string(v))
		} else if len(myStack) == 0 || parenMap[myStack.Pop()] != string(v) {
			return false
		}
	}
	return len(myStack) == 0
}
