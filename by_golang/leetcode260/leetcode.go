package leetcode260

import "fmt"

func singleNumber(nums []int) []int {
	xor := 0
	for _, v := range nums {
		xor ^= v
	}

	lb := xor & -xor

	a, b := 0, 0
	for _, v := range nums {
		if v&lb == 0 {
			a ^= v
		} else {
			b ^= v
		}
	}
	return []int{a, b}
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := singleNumber([]int{2, 2, 1})
	fmt.Printf("is anagram: %v", a)
}
