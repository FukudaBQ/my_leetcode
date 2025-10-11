package leetcode80

import "fmt"

func removeDuplicates(nums []int) int {
	actualpt := len(nums) - 1
	curpt := len(nums) - 1
	for curpt >= 0 {
		if curpt >= 2 && nums[curpt] != nums[curpt-2] {
			nums[actualpt] = nums[curpt]
			curpt--
			actualpt--
		} else if curpt < 2 {
			nums[actualpt] = nums[curpt]
			curpt--
			actualpt--
		} else {
			curpt--
		}
	}
	res := len(nums) - actualpt - 1
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	for i, j := 0, res-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := []int{-1, -1, 1, 1, 1, 1, 2, 3, 3}
	ans := removeDuplicates(a)
	fmt.Printf("ans is: %v\nans is: %v", a, ans)
}
