package leetcode88

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	pt1 := 0
	pt2 := 0
	for i := 0; i < m+n; i++ {
		pt1Num := 1000000001
		if pt1 < m {
			pt1Num = nums1[pt1]
		}
		pt2Num := 1000000001
		if pt2 < n {
			pt2Num = nums2[pt2]
		}
		if pt1Num > pt2Num {
			nums1[(m+i)%len(nums1)] = pt2Num
			pt2++
		} else {
			nums1[(m+i)%len(nums1)] = pt1Num
			pt1++
		}
	}
	for i, j := 0, m+n-1; i < j; i, j = i+1, j-1 {
		nums1[i], nums1[j] = nums1[j], nums1[i]
	}
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		nums1[i], nums1[j] = nums1[j], nums1[i]
	}
	for i, j := n, m+n-1; i < j; i, j = i+1, j-1 {
		nums1[i], nums1[j] = nums1[j], nums1[i]
	}
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := []int{1, 2, 3, 0, 0, 0}
	merge(a, 3, []int{2, 5, 6}, 3)
	fmt.Printf("is anagram: %v", a)
}
