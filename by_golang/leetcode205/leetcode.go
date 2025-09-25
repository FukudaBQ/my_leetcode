package leetcode205

import "fmt"

func isIsomorphic(s string, t string) bool {
	sr := []rune(s)
	tr := []rune(t)
	mapper := make(map[rune]rune)
	revMapper := make(map[rune]rune)
	for i, j := range sr {
		if k, ok := mapper[j]; ok {
			if k != tr[i] {
				return false
			}
		} else {
			mapper[j] = tr[i]
		}
	}

	for i, j := range tr {
		if k, ok := revMapper[j]; ok {
			if k != sr[i] {
				return false
			}
		} else {
			revMapper[j] = sr[i]
		}
	}
	return true
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := isIsomorphic("badc", "baba")
	fmt.Printf("can map: %v", a)
}
