package leetcode292

import "fmt"

func canWinNim(n int) bool {
	return !(n%4 == 0)
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := canWinNim(4)
	fmt.Printf("ans is: %v", ans)
}
