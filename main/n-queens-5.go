package main

import "fmt"

func main() {
	fmt.Println(solveNQueens(4))
}
//设置3个集合 分别存放当前列 撇列 捺列不能方的位置
func solveNQueens(n int) [][]string {
    r := 1
	s,col,pie,na ,res:= make([]string,0),make(map[int]int) ,make(map[int]int),make(map[int]int),make([][]string,0)
	DFS(n,r,s,col,pie,na,&res)
	return res
}

func DFS(n int,r int,s []string,col map[int]int,pie map[int]int,na map[int]int,res *[][]string)  {
	if n<=3 {
		*res = nil
	}
	if n<r {
		*res = append(*res, s)
		return
	}
	for i := 1 ;i<=n;i++  {
		_,ok1 := col[i]
		_,ok2 := pie[i-r]
		_,ok3 := na[i+r]
		if ok2 || ok1||ok3{
			continue
		}
		var str string
		for j:=1;j<=n;j++  {
			if j==i {
				str = str+"Q"
			}else {
				str = str + "."
			}
		}
		col[i] = 1
		pie[i-r] = 1
		na[i+r] = 1
		DFS(n,r+1,append(s,str),col,pie,na,res)
		delete(col, i)
		delete(pie,i-r)
		delete(na,i+r)
	}
}