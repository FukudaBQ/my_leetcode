package leetcode74

import "fmt"

func searchMatrix(matrix [][]int, target int) bool {
	first, last := 0, len(matrix)*len(matrix[0])-1
	for first <= last {
		mid := (first + last) / 2
		row := mid / len(matrix[0])
		col := mid % len(matrix[0])
		fmt.Printf("mid: %d, row: %d, col: %d\n", mid, row, col)
		if matrix[row][col] == target {
			return true
		}
		if matrix[row][col] > target {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return false
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	//a := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 6)
	a := searchMatrix([][]int{{1}}, 6)
	fmt.Printf("is anagram: %v", a)
}
