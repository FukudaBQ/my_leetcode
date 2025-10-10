package leetcode1124

import "fmt"

func longestWPI(hours []int) int {
	prefixSum := make([]int, len(hours)+1)
	prefixSumMap := make(map[int]int)
	prefixSumMap[0] = 0
	maxL := 0
	for i, v := range hours {
		if v > 8 {
			prefixSum[i+1] = prefixSum[i] + 1
		} else {
			prefixSum[i+1] = prefixSum[i] - 1
		}

		if _, ok := prefixSumMap[prefixSum[i+1]]; !ok {
			prefixSumMap[prefixSum[i+1]] = i + 1
		}

		if prefixSum[i+1] > 0 {
			if i+1 > maxL {
				maxL = i + 1
			}
			continue
		}

		if idx, ok := prefixSumMap[prefixSum[i+1]-1]; ok {
			if i+1-idx > maxL {
				maxL = i + 1 - idx
			}
		}
	}
	return maxL
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := longestWPI([]int{6, 6, 9})
	fmt.Printf("is anagram: %v", a)
}
