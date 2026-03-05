package leetcode40

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var res [][]int
	var path []int

	sort.Ints(candidates)

	var backtrack func(start int, remain int)
	backtrack = func(start int, remain int) {
		if remain == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}

		for i := start; i < len(candidates); i++ {
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			if candidates[i] > remain {
				break
			}

			path = append(path, candidates[i])
			backtrack(i+1, remain-candidates[i])
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target)
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := combinationSum2([]int{2, 3, 6, 7}, 7)
	fmt.Printf("water is: %v", a)
}
