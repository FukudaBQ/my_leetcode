package leetcode27

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	j := len(nums) - 1
	for i < j+1 {
		if nums[i] == val {
			nums[i] = nums[j]
			nums[j] = val
			j--
		} else {
			i++
		}
	}
	return j + 1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := removeElement([]int{2, 3, 2}, 2)
	print(a)
}
