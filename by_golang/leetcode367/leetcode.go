package leetcode367

import "fmt"

func isPerfectSquare(num int) bool {
	first, last := 0, num
	for first <= last {
		mid := (first + last) / 2
		if mid*mid == num {
			return true
		}
		if mid*mid < num && (mid+1)*(mid+1) > num {
			return false
		}
		if mid*mid > num {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return false
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := isPerfectSquare(16)
	fmt.Printf("is anagram: %v", a)
}
