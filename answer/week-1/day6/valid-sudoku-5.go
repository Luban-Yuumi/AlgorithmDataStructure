package day6

func isValidSudoku(board [][]byte) bool {
	if len(board) < 9 {
		return false
	}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' && !sudokuHelper(i, j, board) {
				return false
			}
		}
	}

	return true
}

func sudokuHelper(row int, col int, board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if row != i && board[i][col] == board[row][col] {
			return false
		}
		if col != i && board[row][i] == board[row][col] {
			return false
		}

		if board[(row/3)*3+i/3][(col/3)*3+i%3] == board[row][col] && ((row/3)*3+i/3 != row || (col/3)*3+i%3 != col) {
			return false
		}

	}
	return true
}
