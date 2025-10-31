package leetcode704

import "fmt"

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	i := len(nums) / 2
	if nums[i] == target {
		return i
	}
	if nums[i] > target {
		return search(nums[:i], target)
	}
	if sub := search(nums[i+1:], target); sub != -1 {
		return i + 1 + sub
	}
	return -1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := search([]int{-1, 0, 3, 5, 10, 12}, 9)
	fmt.Printf("is anagram: %v", a)
}
