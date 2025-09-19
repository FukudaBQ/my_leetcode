package leetcode283

import (
	"fmt"
)

func moveZeroes(nums []int) {
	pointer := 0
	zeros := 0
	for ; pointer+zeros < len(nums); pointer++ {
		if nums[pointer] == 0 {
			nums := append(nums[:pointer], nums[pointer+1:]...)
			nums = append(nums, 0)
			pointer--
			zeros++
		}
	}
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	print(nums)
}
