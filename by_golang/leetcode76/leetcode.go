package leetcode76

import "fmt"

func minWindow(s string, t string) string {
	sr, tr := []rune(s), []rune(t)
	need := make(map[rune]int)
	for _, ch := range tr {
		need[ch]++
	}
	trMap := make(map[rune]int)
	for ch, v := range need {
		trMap[ch] = v
	}

	pt0, pt1 := 0, 0
	find := 0
	minL, minpt0, minpt1 := len(sr)+1, -1, -1

	for pt1 < len(sr) {
		if find < len(need) {
			ch := sr[pt1]
			if lack, ok := trMap[ch]; ok {
				trMap[ch]--
				if lack == 1 {
					find++
				}
			}
			pt1++
		} else {
			if curLen := pt1 - pt0; curLen < minL {
				minL, minpt0, minpt1 = curLen, pt0, pt1
			}
			// 左指针收缩
			ch := sr[pt0]
			if lack, ok := trMap[ch]; ok {
				if lack == 0 {
					find--
				}
				trMap[ch]++
			}
			pt0++
		}
	}

	for find == len(need) && pt0 < len(sr) {
		if curLen := pt1 - pt0; curLen < minL {
			minL, minpt0, minpt1 = curLen, pt0, pt1
		}
		ch := sr[pt0]
		if lack, ok := trMap[ch]; ok {
			if lack == 0 {
				find--
			}
			trMap[ch]++
		}
		pt0++
	}

	if minpt0 == -1 {
		return ""
	}
	return string(sr[minpt0:minpt1])
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := minWindow("ADOBECODEBANC", "ABC")
	fmt.Printf("is anagram: %v", a)
}
