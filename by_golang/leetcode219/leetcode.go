package leetcode219

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {
	dup := make(map[int]int)
	i := 0
	j := 0
	for j < k && j < len(nums) {
		if _, ok := dup[nums[j]]; ok {
			return true
		}
		dup[nums[j]] = dup[nums[j]] + 1
		j++
	}
	for j < len(nums) {
		if _, ok := dup[nums[j]]; ok {
			return true
		}
		dup[nums[j]] = 0
		dup[nums[i]] = dup[nums[i]] - 1
		if dup[nums[i]] <= 0 {
			delete(dup, nums[i])
		}
		i++
		j++
	}
	return false
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := containsNearbyDuplicate([]int{1, 0, 1, 1}, 1)
	fmt.Printf("contains duplicate: %v", a)
}
