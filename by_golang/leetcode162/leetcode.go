package leetcode162

import "fmt"

func findPeakElement(nums []int) int {
	first, last := 0, len(nums)-1
	for first <= last {
		mid := (first + last) / 2
		prev := nums[mid] - 1
		if mid != 0 {
			prev = nums[mid-1]
		}
		next := nums[mid] - 1
		if mid != len(nums)-1 {
			next = nums[mid+1]
		}
		if nums[mid] > prev && nums[mid] > next {
			return mid
		}
		if nums[mid] > next {
			last = mid - 1
		} else {
			first = mid + 1
		}
	}
	return -1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := findPeakElement([]int{1, 2, 3, 1})
	fmt.Printf("is anagram: %v", a)
}
