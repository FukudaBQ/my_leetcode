package leetcode567

import "fmt"

func checkInclusion(s1 string, s2 string) bool {
	if len(s2) < len(s1) {
		return false
	}

	need := make(map[rune]int)
	for _, ch := range s1 {
		need[ch]++
	}

	diff := make(map[rune]int)
	for k, v := range need {
		diff[k] = v
	}

	var match int
	var res []int

	for right, left := 0, 0; right < len(s2); right++ {
		ch := rune(s2[right])

		if diff[ch] > 0 {
			diff[ch]--
			if diff[ch] == 0 {
				match++
			}
		} else {
			diff[ch]--
		}

		for right-left+1 > len(s1) {
			leftCh := rune(s2[left])
			if diff[leftCh] == 0 {
				match--
			}
			diff[leftCh]++
			left++
		}

		if match == len(need) && right-left+1 == len(s1) {
			res = append(res, left)
		}
	}
	return len(res) > 0
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := checkInclusion("cbaebabacd", "abc")
	fmt.Printf("is anagram: %v", a)
}
