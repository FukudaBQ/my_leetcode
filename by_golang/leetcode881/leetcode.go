package leetcode881

import (
	"fmt"
	"sort"
)

func numRescueBoats(people []int, limit int) int {
	sort.Ints(people)
	head := 0
	tail := len(people) - 1
	res := 0
	for head < tail {
		front := people[head]
		back := people[tail]
		if front+back <= limit {
			res++
			head++
			tail--
		} else {
			res++
			tail--
		}
	}
	if head == tail {
		res++
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := numRescueBoats([]int{3, 2, 2, 1}, 3)
	fmt.Printf("is anagram: %v", a)
}
