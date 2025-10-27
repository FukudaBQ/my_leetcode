package leetcode55

import "fmt"

func canJump(nums []int) bool {
	maxReach := 0
	for i := 0; i < len(nums); i++ {
		if maxReach < i {
			return false
		}
		if i+nums[i] > maxReach {
			maxReach = i + nums[i]
		}
	}
	return true
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := canJump([]int{2, 3, 1, 1, 4})
	fmt.Printf("is anagram: %v", a)
}
