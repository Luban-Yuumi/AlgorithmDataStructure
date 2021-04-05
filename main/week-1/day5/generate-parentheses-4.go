package day5

//方法1 暴力枚举 时间复杂度（o(2n)）
//方法2 剪枝
//（1）左括号数目 右括号数目都小于n
// (2) 左括号数目总是要大于等于右括号数目的
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
