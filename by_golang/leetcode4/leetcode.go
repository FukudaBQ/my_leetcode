package leetcode4

import (
	"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	len1, len2 := len(nums1), len(nums2)
	med := (len1 + len2) / 2
	
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := findMedianSortedArrays([]int{1, 2}, []int{3, 4})
	fmt.Println(a)
}
