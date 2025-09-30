package leetcode49

import (
	"fmt"
)

func groupAnagrams(strs []string) [][]string {
	strmap := make(map[[26]int][]string)
	for _, i := range strs {
		runeStr := [26]int{}
		for _, runeval := range i {
			runeStr[runeval-'a']++
		}
		if _, ok := strmap[runeStr]; ok {
			slice := strmap[runeStr]
			slice = append(slice, i)
			strmap[runeStr] = slice
		} else {
			strmap[runeStr] = []string{i}
		}
	}
	var result [][]string
	for _, v := range strmap {
		result = append(result, v)
	}
	return result
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Printf("is anagram: %v", a)
}
