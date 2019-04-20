package main

import "fmt"

func solveSudoku(board [][]byte)  {
	solveSudoku1(board)
}

func solveSudoku1(board [][]byte) bool {
	var charSet = []byte{'1','2','3','4','5','6','7','8','9'}
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for _,v := range charSet {
					if  isValid(i,j,board,v){
						board[i][j] = v
						if (solveSudoku1(board)){
							return true
						}else {
							board[i][j] = '.'
						}
					}
				}
				return false
			}
		}
	}
	return true
}
func isValid(row int,col int,board [][]byte,v byte)bool  {
	for i:=0;i<9 ;i++  {
		if board[i][col]!='.'&&board[i][col]==v {
			return false
		}
		if board[row][i]!='.'&&board[row][i]==v {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3]!= '.' && board[(row/3)*3+i/3][(col/3)*3+i%3] == v  {
			return false
		}
	}
	return true
}