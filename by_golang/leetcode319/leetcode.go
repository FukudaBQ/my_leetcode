package leetcode319

import (
	"fmt"
)

func bulbSwitch(n int) int {
	ans := 0
	for i := 1; i*i <= n; i++ {
		ans++

	}
	return ans
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := bulbSwitch(3)
	fmt.Printf("ans is: %v", ans)
}
