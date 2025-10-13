package leetcode9

import (
	"fmt"
)

func isPalindrome(x int) bool {
	y := x
	if x < 0 {
		return false
	} else if x < 10 {
		return true
	}
	reverse := 0
	for x != 0 {
		reminder := x % 10
		reverse = reverse*10 + reminder
		x = x / 10
	}
	return reverse == y
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := isPalindrome(1213)
	fmt.Printf("ans is: %v", ans)
}
