package leetcode485

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	fp := 0
	lp := 0
	ma := 0
	nums = append(nums, 0)
	nums = append([]int{0}, nums...)
	for ; lp < len(nums); lp++ {
		if nums[lp] == 1 {
		} else {
			if ma < lp-fp {
				ma = lp - fp - 1
			}
			fp = lp
		}
	}
	return ma
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := findMaxConsecutiveOnes([]int{1, 1, 0, 1, 1, 1})
	print(a)
}
