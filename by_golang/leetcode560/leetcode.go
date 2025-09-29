package leetcode560

import "fmt"

func subarraySum(nums []int, k int) int {
	prefixSum := make([]int, len(nums)+1, len(nums)+1)
	prefixSum[0] = 0
	prefixSumMap := make(map[int]int, len(nums)+1)
	prefixSumMap[0] = 1
	count := 0
	for i, j := range nums {
		prefixSum[i+1] = prefixSum[i] + j
		key := prefixSum[i+1] - k
		if _, ok := prefixSumMap[key]; ok {
			count += prefixSumMap[key]
		}
		if _, ok := prefixSumMap[prefixSum[i+1]]; ok {
			prefixSumMap[prefixSum[i+1]]++
		} else {
			prefixSumMap[prefixSum[i+1]] = 1
		}
	}
	return count
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := subarraySum([]int{-1, -1, 1}, 0)
	fmt.Printf("is anagram: %v", a)
}
