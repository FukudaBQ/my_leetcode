package leetcode200

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	stack := make([]int, 0, rows*cols)
	traversed := make(map[int]bool)
	result := 0

	posToInt := func(i, j int) int {
		return i*cols + j
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			pos := posToInt(i, j)
			if !traversed[pos] && grid[i][j] == '1' {
				result++

				stack = stack[:0]
				stack = append(stack, pos)
				traversed[pos] = true

				for len(stack) > 0 {
					current := stack[len(stack)-1]
					stack = stack[:len(stack)-1]

					ci, cj := current/cols, current%cols

					directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
					for _, dir := range directions {
						ni, nj := ci+dir[0], cj+dir[1]

						if ni < 0 || ni >= rows || nj < 0 || nj >= cols {
							continue
						}

						neighborPos := posToInt(ni, nj)
						if !traversed[neighborPos] && grid[ni][nj] == '1' {
							traversed[neighborPos] = true
							stack = append(stack, neighborPos)
						}
					}
				}
			}
		}
	}
	return result
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := numIslands([][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0}})
	fmt.Printf("Result is: %v", a)
}
