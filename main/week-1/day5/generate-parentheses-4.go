package day5

//方法1 暴力枚举 时间复杂度（o(2n)）
//方法2 剪枝
//（1）左括号数目 右括号数目都小于n
// (2) 左括号数目总是要大于等于右括号数目的
func generateParenthesis(n int) []string {
	var res = new([]string)
	helper(n, 0, 0, "", res)
	return *res
}

func helper(n int, left int, right int, s string, res *[]string) {
	if right == n {
		*res = append(*res, s)
	}
	if left < n {
		helper(n, left+1, right, s+"(", res)
	}
	if right < left {
		helper(n, left, right+1, s+")", res)
	}
}
