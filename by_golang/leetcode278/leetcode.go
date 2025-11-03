package leetcode278

import "fmt"

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int, isBadVersion func(int) bool) int {
	first := 0
	last := n
	for first <= last {
		mid := (first + last) / 2
		if isBadVersion(mid) {
			if !isBadVersion(mid-1) || mid == 1 {
				return mid
			}
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return -1
}

func buildIsBadVersion(bad int) func(int) bool {
	return func(v int) bool {
		return v >= bad
	}
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := firstBadVersion(1, buildIsBadVersion(1))
	fmt.Printf("is anagram: %v", a)
}
