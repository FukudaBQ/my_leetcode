package leetcode51

import "fmt"

func solveNQueens(n int) [][]string {
	board := make([][]byte, n)
	for i := range board {
		board[i] = make([]byte, n)
		for j := range board[i] {
			board[i][j] = '.'
		}
	}

	var res [][]string

	var backtrack func(row int)
	backtrack = func(row int) {
		if row == n {
			temp := make([]string, n)
			for i := 0; i < n; i++ {
				temp[i] = string(board[i])
			}
			res = append(res, temp)
			return
		}

		for col := 0; col < n; col++ {
			board[row][col] = 'Q'

			if check(board, n) {
				backtrack(row + 1)
			}

			board[row][col] = '.'
		}
	}

	backtrack(0)
	return res
}

func check(board [][]byte, n int) bool {
	var queens [][2]int
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == 'Q' {
				queens = append(queens, [2]int{i, j})
			}
		}
	}

	for i := 0; i < len(queens); i++ {
		for j := i + 1; j < len(queens); j++ {
			r1, c1 := queens[i][0], queens[i][1]
			r2, c2 := queens[j][0], queens[j][1]

			if r1 == r2 || c1 == c2 || abs(r1-r2) == abs(c1-c2) {
				return false
			}
		}
	}

	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := solveNQueens(4)
	fmt.Printf("result is: %v", a)
}
