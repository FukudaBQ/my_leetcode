package leetcode3

import "fmt"

func lengthOfLongestSubstring(s string) int {
	hashtable := make(map[rune]int)
	rs := []rune(s)
	pt0, pt1 := 0, 0
	res := 0
	for pt1 < len(rs) {
		char := rs[pt1]
		if lastPos, ok := hashtable[char]; ok {
			if lastPos+1 > pt0 {
				pt0 = lastPos + 1
			}
		}
		hashtable[char] = pt1
		length := pt1 - pt0 + 1
		if length > res {
			res = length
		}
		pt1++
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := lengthOfLongestSubstring("abba")
	fmt.Printf("is anagram: %v", a)
}
