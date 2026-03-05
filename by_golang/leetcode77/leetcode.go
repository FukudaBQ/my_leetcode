package leetcode77

import (
	"fmt"
)

func combine(n int, k int) [][]int {
	res := make([][]int, 0)
	path := make([]int, 0)

	var backtrack func(start int)
	backtrack = func(start int) {
		if len(path)+(n-start+1) < k {
			return
		}

		if len(path) == k {
			comb := make([]int, k)
			copy(comb, path)
			res = append(res, comb)
			return
		}

		for i := start; i <= n; i++ {
			path = append(path, i)
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(1)
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := combine(4, 2)
	fmt.Printf("water is: %v", a)
}
