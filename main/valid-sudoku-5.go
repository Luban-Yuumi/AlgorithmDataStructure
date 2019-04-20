package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	fmt.Println(len(board))
	if len(board) <9  {
		return false
	}
	for i:=0;i<9 ;i++  {
		for j:=0; j<9;j++{
			if board[i][j]!='.' {
				if !isValid(i,j,board) {
					return false
				}
			}
		}
	}
	return true
}

func isValid(row int,col int,board [][]byte)bool  {
	for i:=0;i<9 ;i++  {
		if i!=row&&board[i][col]==board[row][col] {
			return false
		}
		if i!=col&&board[row][i]==board[row][col] {
			return false
		}
		if board[(row/3)*3+i/3][(col/3)*3+i%3] == board[row][col] &&(((row/3)*3+i/3)!=row||((col/3)*3+i%3)!=col) {
			return false
		}
	}
	return true
}