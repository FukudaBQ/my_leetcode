package leetcode387

import "fmt"

func firstUniqChar(s string) int {
	set := make(map[rune]int)
	for _, r := range s {
		if _, ok := set[r]; ok {
			set[r]++
		} else {
			set[r] = 1
		}
	}
	for i, r := range s {
		if set[r] == 1 {
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

	a := firstUniqChar("leetcode")
	fmt.Printf("is anagram: %v", a)
}
