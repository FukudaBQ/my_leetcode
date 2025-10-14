package leetcode455

import (
	"fmt"
	"sort"
)

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	res := 0
	i := 0
	j := 0
	for i < len(g) && j < len(s) {
		if s[j] >= g[i] {
			i++
			j++
			res++
		} else {
			j++
		}
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8})
	fmt.Printf("ans is: %v", ans)
}
