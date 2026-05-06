package leetcode37

import "fmt"

func solveSudoku(board [][]byte) {
	backtrack(board)
}

func backtrack(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for num := byte('1'); num <= '9'; num++ {
					if isValid(board, i, j, num) {
						board[i][j] = num
						if backtrack(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, num byte) bool {
	for j := 0; j < 9; j++ {
		if board[row][j] == num {
			return false
		}
	}

	for i := 0; i < 9; i++ {
		if board[i][col] == num {
			return false
		}
	}

	boxRow, boxCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if board[boxRow+i][boxCol+j] == num {
				return false
			}
		}
	}

	return true
}

func Leetcode() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}

	solveSudoku(board)

	// 打印结果
	for _, row := range board {
		fmt.Println(string(row))
	}
}
