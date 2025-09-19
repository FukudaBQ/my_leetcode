package leetcode26

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	current := -101
	dup := 0
	pointer := 0
	for dup+pointer < len(nums) {
		if nums[pointer] > current {
			current = nums[pointer]
			pointer++
		} else {
			dup++
			cur := nums[pointer]
			tmp := append(nums[:pointer], nums[pointer+1:]...)
			nums = append(tmp, cur)
		}
	}
	nums = nums[:len(nums)-dup]
	return len(nums)
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := removeDuplicates([]int{1, 1, 2})
	print(a)
}
