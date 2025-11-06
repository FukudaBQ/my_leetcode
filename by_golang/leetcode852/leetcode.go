package leetcode852

import "fmt"

func peakIndexInMountainArray(arr []int) int {
	l, r := 1, len(arr)-2
	for l <= r {
		m := (l + r) / 2
		if arr[m] > arr[m-1] && arr[m] > arr[m+1] {
			return m
		}
		if arr[m] < arr[m+1] {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := peakIndexInMountainArray([]int{1, 2, 3, 1})
	fmt.Printf("is anagram: %v", a)
}
