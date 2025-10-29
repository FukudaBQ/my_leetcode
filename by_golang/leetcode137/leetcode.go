package leetcode137

import "fmt"

func singleNumber(nums []int) int {
	var once, twice int
	for _, v := range nums {
		twice = twice | once&v
		once ^= v
		mask := ^(once & twice)
		once &= mask
		twice &= mask
	}
	return once
}
func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := singleNumber([]int{2, 2, 1})
	fmt.Printf("is anagram: %v", a)
}
