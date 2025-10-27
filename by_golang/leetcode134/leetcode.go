package leetcode134

import "fmt"

func canCompleteCircuit(gas []int, cost []int) int {
	pt := 0
	res := 0
	ans := -1
	for pt < len(gas) {
		if gas[pt]-cost[pt]+res >= 0 {
			res += gas[pt] - cost[pt]
			if ans == -1 {
				ans = pt
			}
		} else {
			res = 0
			ans = -1
		}
		pt++
	}
	gasR := 0
	if ans == -1 {
		return ans
	}
	for i := 0; i < len(gas); i++ {
		if gasR+gas[(ans+i)%len(gas)]-cost[(ans+i)%len(gas)] >= 0 {
			gasR += gas[(ans+i)%len(gas)] - cost[(ans+i)%len(gas)]
		} else {
			return -1
		}
	}
	return ans
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := canCompleteCircuit([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2})
	fmt.Printf("is anagram: %v", a)
}
