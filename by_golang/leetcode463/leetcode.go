package leetcode463

import "fmt"

func islandPerimeter(grid [][]int) int {
	rows := len(grid)
	if rows == 0 {
		return 0
	}
	cols := len(grid[0])

	stack := make([]int, 0, rows*cols)
	visited := make(map[int]bool)
	posToInt := func(i, j int) int { return i*cols + j }

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 {
				stack = append(stack, posToInt(i, j))
				visited[posToInt(i, j)] = true
				perimeter := 0

				for len(stack) > 0 {
					cur := stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					ci, cj := cur/cols, cur%cols

					cellPeri := 4
					for _, d := range [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}} {
						ni, nj := ci+d[0], cj+d[1]
						if ni >= 0 && ni < rows && nj >= 0 && nj < cols && grid[ni][nj] == 1 {
							cellPeri--
							pos := posToInt(ni, nj)
							if !visited[pos] {
								visited[pos] = true
								stack = append(stack, pos)
							}
						}
					}
					perimeter += cellPeri
				}
				return perimeter
			}
		}
	}
	return 0
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := islandPerimeter([][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}})
	fmt.Printf("Result is: %v", a)
}
