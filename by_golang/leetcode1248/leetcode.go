package leetcode1248

import "fmt"

func numberOfSubarrays(nums []int, k int) int {
	prefixOdd := make([]int, len(nums)+1)
	prefixOdd[0] = 0
	prefixOddMap := make(map[int]int, len(nums)+1)
	count := 0
	prefixOddMap[0] = 1
	for i, j := range nums {
		if j%2 == 1 {
			prefixOdd[i+1] = prefixOdd[i] + 1
		} else {
			prefixOdd[i+1] = prefixOdd[i]
		}
		key := prefixOdd[i+1] - k
		if _, ok := prefixOddMap[key]; ok {
			count += prefixOddMap[key]
		}
		if _, ok := prefixOddMap[prefixOdd[i+1]]; ok {
			prefixOddMap[prefixOdd[i+1]]++
		} else {
			prefixOddMap[prefixOdd[i+1]] = 1
		}
	}
	return count
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := numberOfSubarrays([]int{1, 1, 2, 1, 1}, 3)
	fmt.Printf("is anagram: %v", a)
}
