package leetcode56

import (
	"fmt"
	"sort"
)

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0, len(intervals))
	curStart := intervals[0][0]
	curEnd := intervals[0][1]
	for i := 0; i < len(intervals); i++ {
		cur := intervals[i]
		if cur[0] <= curEnd {
			if cur[1] > curEnd {
				curEnd = cur[1]
			}
		} else {
			res = append(res, []int{curStart, curEnd})
			curStart = cur[0]
			curEnd = cur[1]
		}
	}
	res = append(res, []int{curStart, curEnd})
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := merge([][]int{
		{1, 3},
		{2, 6},
		{8, 10},
		{15, 18}})
	fmt.Printf("Result is: %v", a)
}
