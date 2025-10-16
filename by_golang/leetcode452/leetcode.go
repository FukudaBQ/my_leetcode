package leetcode452

import (
	"fmt"
	"sort"
)

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	arrow := 1
	end := points[0][1]
	for i := 0; i < len(points); i++ {
		if points[i][0] <= end {
			if points[i][1] <= end {
				end = points[i][1]
			}
		} else {
			arrow++
			end = points[i][1]
		}
	}
	return arrow
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})
	fmt.Printf("ans is: %v", ans)
}
