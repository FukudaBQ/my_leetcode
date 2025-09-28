package leetcode409

import "fmt"

func longestPalindrome(s string) int {
	set := make(map[rune]int)
	for _, j := range s {
		if _, ok := set[j]; ok {
			set[j]++
		} else {
			set[j] = 1
		}
	}
	res := 0
	for _, v := range set {
		res += v / 2
	}
	res *= 2
	if res < len(s) {
		res++
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := longestPalindrome("abccccdd")
	fmt.Printf("is anagram: %v", a)
}
