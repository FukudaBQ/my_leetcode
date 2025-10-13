package leetcode680

import (
	"fmt"
)

func validPalindrome(s string) bool {
	sc := []rune(s)
	return checkPalindrome(sc, true)
}

func checkPalindrome(s []rune, tolerance bool) bool {
	if len(s) < 2 {
		return true
	}
	pt1 := s[0]
	pt2 := s[len(s)-1]
	if pt1 != pt2 {
		if !tolerance {
			return false
		}
		return checkPalindrome(s[1:], false) || checkPalindrome(s[:len(s)-1], false)
	}
	return checkPalindrome(s[1:len(s)-1], tolerance)
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := validPalindrome("abc")
	fmt.Printf("ans is: %v", ans)
}
