package leetcode402

import (
	"fmt"
	"strings"
)

func removeKdigits(num string, k int) string {
	stack := make([]rune, 0, len(num))
	for _, d := range num {
		for k > 0 && len(stack) > 0 && stack[len(stack)-1] > d {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, d)
	}

	if k > 0 {
		stack = stack[:len(stack)-k]
	}

	trimmed := strings.TrimLeftFunc(string(stack), func(r rune) bool {
		return r == '0'
	})
	if trimmed == "" {
		return "0"
	}
	return trimmed
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := removeKdigits("1432219", 3)
	fmt.Printf("ans is: %v", ans)
}
