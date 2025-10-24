package leetcode11

import "fmt"

func maxArea(height []int) int {
	pt0, pt1 := 0, len(height)-1
	maxA := 0
	for pt0 < pt1 {
		l := 0
		if height[pt0] < height[pt1] {
			l = height[pt0]
		} else {
			l = height[pt1]
		}
		if l*(pt1-pt0) > maxA {
			maxA = l * (pt1 - pt0)
		}
		if height[pt0] < height[pt1] {
			pt0++
		} else {
			pt1--
		}
	}
	return maxA
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	fmt.Printf("ans is: %v", ans)
}
