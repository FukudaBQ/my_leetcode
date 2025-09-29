package leetcode724

import "fmt"

func pivotIndex(nums []int) int {
	sum := make([]int, len(nums)+1)
	for i, j := range nums {
		sum[i+1] = sum[i] + j
	}
	for i, j := range sum {
		rightSum := 0
		if i+1 < len(sum) {
			rightSum = sum[len(sum)-1] - sum[i+1]
		} else {
			return -1
		}
		if j == rightSum {
			return i
		}
	}
	return -1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := pivotIndex([]int{-1, -1, -1, 1, 1, 1})
	fmt.Printf("is anagram: %v", a)
}
