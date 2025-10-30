package leetcode8

import (
	"fmt"
	"math"
	"unicode"
)

func myAtoi(s string) int {
	sr := []rune(s)
	pt := 0
	sign := 1
	res := 0
	for pt < len(sr) && sr[pt] == ' ' {
		pt++
	}
	if pt == len(sr) {
		return 0
	}
	if sr[pt] == '-' {
		sign = -1
		pt++
	} else if sr[pt] == '+' {
		sign = 1
		pt++
	}
	for ; pt < len(sr) && unicode.IsDigit(sr[pt]); pt++ {
		res = res*10 + int(sr[pt]-'0')
		if sign == 1 && res > math.MaxInt32 {
			return math.MaxInt32
		} else if sign == -1 && res > -math.MinInt32 {
			return math.MinInt32
		}
	}
	return sign * res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := myAtoi("42")
	fmt.Printf("is anagram: %v", a)
}
