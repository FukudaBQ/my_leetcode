package leetcode36

import (
	"encoding/json"
	"fmt"
)

func isValidSudoku(board [][]byte) bool {
	rowMap := make(map[int][]int)
	columnMap := make(map[int][]int)
	nineMap := make(map[int][]int)
	for row, rValue := range board {
		for column, cValue := range rValue {
			var columns []int
			if _, ok := columnMap[column]; !ok {
				columns = make([]int, 10)
			} else {
				columns = columnMap[column]
			}

			var rows []int
			if _, ok := rowMap[row]; !ok {
				rows = make([]int, 10)
			} else {
				rows = rowMap[row]
			}

			var grid []int
			gridKey := row/3*3 + column/3
			if _, ok := nineMap[gridKey]; !ok {
				grid = make([]int, 10)
			} else {
				grid = nineMap[gridKey]
			}

			if cValue == '.' {
				columns[9]++
				rows[9]++
				grid[9]++
			} else {
				columns[cValue-'1']++
				rows[cValue-'1']++
				grid[cValue-'1']++
			}

			rowMap[row] = rows
			columnMap[column] = columns
			nineMap[gridKey] = grid
		}
	}

	for _, value := range rowMap {
		sum := 0
		for i, ins := range value {
			if ins > 1 && i != 9 {
				return false
			}
			sum += ins
		}
		if sum > 9 {
			return false
		}
	}

	for _, value := range columnMap {
		sum := 0
		for i, ins := range value {
			if ins > 1 && i != 9 {
				return false
			}
			sum += ins
		}
		if sum > 9 {
			return false
		}
	}

	for _, value := range nineMap {
		sum := 0
		for i, ins := range value {
			if ins > 1 && i != 9 {
				return false
			}
			sum += ins
		}
		if sum > 9 {
			return false
		}
	}
	return true
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	input, err := strToBoard("[[\".\",\".\",\".\",\".\",\".\",\".\",\".\",\".\",\".\"],[\".\",\".\",\".\",\".\",\".\",\".\",\".\",\".\",\".\"],[\".\",\"9\",\".\",\".\",\".\",\".\",\".\",\".\",\"1\"],[\"8\",\".\",\".\",\".\",\".\",\".\",\".\",\".\",\".\"],[\".\",\"9\",\"9\",\"3\",\"5\",\"7\",\".\",\".\",\".\"],[\".\",\".\",\".\",\".\",\".\",\".\",\".\",\"4\",\".\"],[\".\",\".\",\".\",\"8\",\".\",\".\",\".\",\".\",\".\"],[\".\",\"1\",\".\",\".\",\".\",\".\",\"4\",\".\",\"9\"],[\".\",\".\",\".\",\"5\",\".\",\"4\",\".\",\".\",\".\"]]")
	//input, err := strToBoard("[[\"5\",\"3\",\".\",\".\",\"7\",\".\",\".\",\".\",\".\"],[\"6\",\".\",\".\",\"1\",\"9\",\"5\",\".\",\".\",\".\"],[\".\",\"9\",\"8\",\".\",\".\",\".\",\".\",\"6\",\".\"],[\"8\",\".\",\".\",\".\",\"6\",\".\",\".\",\".\",\"3\"],[\"4\",\".\",\".\",\"8\",\".\",\"3\",\".\",\".\",\"1\"],[\"7\",\".\",\".\",\".\",\"2\",\".\",\".\",\".\",\"6\"],[\".\",\"6\",\".\",\".\",\".\",\".\",\"2\",\"8\",\".\"],[\".\",\".\",\".\",\"4\",\"1\",\"9\",\".\",\".\",\"5\"],[\".\",\".\",\".\",\".\",\"8\",\".\",\".\",\"7\",\"9\"]]")
	printBoard(input)
	if err != nil {
		fmt.Println(err)
	}
	a := isValidSudoku(input)
	fmt.Printf("is anagram: %v", a)
}

func strToBoard(in string) ([][]byte, error) {
	var tmp [][]string
	if err := json.Unmarshal([]byte(in), &tmp); err != nil {
		return nil, err
	}

	board := make([][]byte, 9)
	for i := 0; i < 9; i++ {
		board[i] = make([]byte, 9)
		for j := 0; j < 9; j++ {
			if tmp[i][j] == "." {
				board[i][j] = '.'
			} else {
				board[i][j] = tmp[i][j][0]
			}
		}
	}
	return board, nil
}

func printBoard(board [][]byte) {
	fmt.Println("   +-------+-------+-------+")
	for r := 0; r < 9; r++ {
		if r == 0 {
			fmt.Print("   ")
			for c := 0; c < 9; c++ {
				if c%3 == 0 {
					fmt.Print("  ")
				}
				fmt.Printf(" %d ", c+1)
			}
			fmt.Println()
			fmt.Println("   +-------+-------+-------+")
		}

		fmt.Printf(" %d |", r+1)

		for c := 0; c < 9; c++ {
			ch := board[r][c]
			if ch == '.' {
				ch = ' '
			}
			fmt.Printf(" %c ", ch)
			if c%3 == 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()

		if r%3 == 2 {
			fmt.Println("   +-------+-------+-------+")
		}
	}
}
