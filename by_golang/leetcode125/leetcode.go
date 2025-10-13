package leetcode125

import (
	"fmt"
	"unicode"
)

func isPalindrome(s string) bool {
	pt1 := 0
	pt2 := len(s) - 1
	sc := []rune(s)
	for pt1 <= pt2 {
		pt1c := sc[pt1]
		if (!unicode.IsDigit(pt1c)) && !unicode.IsLetter(pt1c) {
			pt1++
			continue
		}
		if unicode.IsLetter(pt1c) {
			pt1c = unicode.ToLower(pt1c)
		}
		pt2c := sc[pt2]
		if (!unicode.IsDigit(pt2c)) && !unicode.IsLetter(pt2c) {
			pt2--
			continue
		}
		if unicode.IsLetter(pt2c) {
			pt2c = unicode.ToLower(pt2c)
		}
		if pt1c != pt2c {
			return false
		}
		pt1++
		pt2--
	}
	return true
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := isPalindrome("A man, a plan, a canal: Panama")
	fmt.Printf("ans is: %v", ans)
}
