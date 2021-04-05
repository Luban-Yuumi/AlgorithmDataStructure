package main

func generateParenthesis(n int) []string {
	var res []string
	helper(n, n, "", &res)
	return res
}

func helper(left int, right int, str string, res *[]string) {
	if right == 0 {
		*res = append(*res, str)
		return
	}

	if left > 0 {
		helper(left-1, right, str+"(", res)
	}
	if right > left {
		helper(left, right-1, str+")", res)
	}
	return
}
