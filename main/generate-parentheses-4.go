package main

//方法1 暴力枚举 时间复杂度（o(2n)）
//方法2 剪枝
//（1）左括号数目 右括号数目都小于n
// (2) 左括号数目总是要大于等于右括号数目的
func generateParenthesis(n int) []string {
	var res = make([]string, 0)
	var left = 0
	var right = 0
	_gen(left, right, &res, "", n)
	return res
}
func _gen(left int, right int, res *[]string, nowString string, n int) {
	if left == n && right == n {
		*res = append(*res, nowString)
	}
	if left < n {
		_gen(left+1, right, res, nowString+"(", n)
	}
	if right < n && right < left {
		_gen(left, right+1, res, nowString+")", n)
	}
}
