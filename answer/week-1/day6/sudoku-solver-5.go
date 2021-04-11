package day6

var charSet = []byte{'1', '2', '3', '4', '5', '6', '7', '8', '9'}

func solveSudoku(board [][]byte) {
	solveSudokuHelper(board)
}

func solveSudokuHelper(board [][]byte) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			if board[row][col] == '.' {
				for i := range charSet {
					if !isValid(row, col, board, charSet[i]) {
						continue
					}
					board[row][col] = charSet[i]
					if solveSudokuHelper(board) {
						return true
					} else {
						board[row][col] = '.'
						continue
					}
				}
				return false
			}
		}
	}

	return true
}

func isValid(row int, col int, board [][]byte, v byte) bool {
	for i := 0; i < 9; i++ {
		if board[row][i] == v || board[i][col] == v || board[(row/3)*3+i/3][(col/3)*3+i%3] == v {
			return false
		}
	}
	return true
}
