package leetcode217

import "fmt"

func containsDuplicate(nums []int) bool {
	hashset := make(map[int]struct{})
	for _, j := range nums {
		if _, ok := hashset[j]; ok {
			return true
		}
		hashset[j] = struct{}{}
	}
	return false
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := containsDuplicate([]int{1, 2, 3, 4})
	fmt.Printf("contains duplicate: %v", a)
}
