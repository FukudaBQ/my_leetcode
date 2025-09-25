package leetcode349

import "fmt"

func intersection(nums1 []int, nums2 []int) []int {
	var short, long []int
	if len(nums1) < len(nums2) {
		short = nums1
		long = nums2
	} else {
		short = nums2
		long = nums1
	}
	dupSet := make(map[int]struct{})
	resMap := make(map[int]struct{})
	for _, j := range short {
		if _, ok := dupSet[j]; !ok {
			dupSet[j] = struct{}{}
		}
	}
	for _, j := range long {
		if _, ok := dupSet[j]; ok {
			if _, ok := resMap[j]; !ok {
				resMap[j] = struct{}{}
			}
		}
	}
	res := make([]int, 0)
	for k, _ := range resMap {
		res = append(res, k)
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := intersection([]int{1, 2, 2, 1}, []int{2, 2})
	fmt.Printf("contains duplicate: %v", a)
}
