package leetcode438

import "fmt"

func findAnagrams(s, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}

	need := make(map[rune]int)
	for _, ch := range p {
		need[ch]++
	}

	diff := make(map[rune]int)
	for k, v := range need {
		diff[k] = v
	}

	var match int
	var res []int

	for right, left := 0, 0; right < len(s); right++ {
		ch := rune(s[right])

		if diff[ch] > 0 {
			diff[ch]--
			if diff[ch] == 0 {
				match++
			}
		} else {
			diff[ch]--
		}

		for right-left+1 > len(p) {
			leftCh := rune(s[left])
			if diff[leftCh] == 0 {
				match--
			}
			diff[leftCh]++
			left++
		}

		if match == len(need) && right-left+1 == len(p) {
			res = append(res, left)
		}
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := findAnagrams("cbaebabacd", "abc")
	fmt.Printf("is anagram: %v", a)
}
