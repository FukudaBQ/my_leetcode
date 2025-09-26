package leetcode242

import "fmt"

func isAnagram(s string, t string) bool {
	set := make(map[rune]int)
	for _, c := range s {
		if _, ok := set[c]; ok {
			set[c]++
		} else {
			set[c] = 1
		}
	}
	for _, c := range t {
		if _, ok := set[c]; ok {
			set[c]--
			if set[c] <= 0 {
				delete(set, c)
			}
		} else {
			return false
		}
	}
	return len(set) == 0
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := isAnagram("anagram", "nagaram")
	fmt.Printf("is anagram: %v", a)
}
