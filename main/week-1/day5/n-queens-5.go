package day5

func solveNQueens(n int) [][]string {
	col, pie, na := make(map[int]struct{}), make(map[int]struct{}), make(map[int]struct{})
	s := make([]string, n)
	res := new([][]string)
	helper(n, 1, col, pie, na, res, s)
	return *res
}

func helper(n int, row int, col map[int]struct{}, pie map[int]struct{}, na map[int]struct{}, res *[][]string, s []string) {
	if row > n {
		tmp := make([]string, n)
		copy(tmp, s)
		*res = append(*res, tmp)
	}

	for j := 1; j <= n; j++ {
		if _, ok := col[j]; ok {
			continue
		}
		if _, ok := pie[j+row]; ok {
			continue
		}
		if _, ok := na[j-row]; ok {
			continue
		}

		col[j] = struct{}{}
		pie[j+row] = struct{}{}
		na[j-row] = struct{}{}
		str := ""
		for k := 1; k <= n; k++ {
			if k == j {
				str += "Q"
			} else {
				str += "."
			}
		}
		s[row-1] = str
		helper(n, row+1, col, pie, na, res, s)
		delete(col, j)
		delete(pie, j+row)
		delete(na, j-row)
	}
}
