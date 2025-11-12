package leetcode1695

import "fmt"

func maximumUniqueSubarray(nums []int) int {
	pt0, pt1, seen := 0, 0, make(map[int]int, len(nums))
	sum := 0
	prefixSum := make([]int, len(nums)+1)
	for i, j := range nums {
		prefixSum[i+1] = prefixSum[i] + j
	}
	for pt1 < len(nums) {
		fmt.Printf("%d %d %v\n", pt0, pt1, seen)
		if lastSeen, ok := seen[nums[pt1]]; ok {
			if lastSeen >= pt0 {
				pt0 = lastSeen + 1
			}
		}
		seen[nums[pt1]] = pt1
		newSum := prefixSum[pt1+1] - prefixSum[pt0]
		if newSum > sum {
			sum = newSum
		}
		pt1++
	}
	return sum
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := maximumUniqueSubarray([]int{4, 2, 4, 5, 6})
	fmt.Printf("is anagram: %v", a)
}
