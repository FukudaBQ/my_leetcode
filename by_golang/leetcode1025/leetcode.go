package leetcode1025

import "fmt"

func divisorGame(n int) bool {
	return n%2 == 0
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := divisorGame(3)
	fmt.Printf("ans is: %v", ans)
}
