package leetcode435

import (
	"fmt"
	"sort"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	remove := 0
	end := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		if intervals[i][0] < end {
			remove++
		} else {
			end = intervals[i][1]
		}
	}
	return remove - 1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := eraseOverlapIntervals([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})
	fmt.Printf("ans is: %v", ans)
}
