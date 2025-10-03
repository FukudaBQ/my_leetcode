package leetcode238

import "fmt"

func productExceptSelf(nums []int) []int {
	products := make([]int, len(nums)+1)
	revProducts := make([]int, len(nums)+1)
	products[0] = 1
	revProducts[len(nums)] = 1
	for i := 0; i < len(nums); i++ {
		products[i+1] = products[i] * nums[i]
		revProducts[len(nums)-1-i] = revProducts[len(nums)-i] * nums[len(nums)-1-i]
	}
	res := make([]int, len(nums))
	for i, _ := range nums {
		res[i] = products[i] * revProducts[i+1]
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := productExceptSelf([]int{1, 2, 3, 4})
	fmt.Printf("is anagram: %v", a)
}
