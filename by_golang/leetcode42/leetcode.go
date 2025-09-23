package leetcode42

import "fmt"

func trap(height []int) int {
	maxLeft := 0
	leftMax := make([]int, 0)
	for _, h := range height {
		leftMax = append(leftMax, maxLeft)
		if h > maxLeft {
			maxLeft = h
		}
	}
	rightMax := make([]int, len(height))
	maxRight := 0
	for i := len(height) - 1; i >= 0; i-- {
		rightMax[i] = maxRight
		if height[i] > maxRight {
			maxRight = height[i]
		}
	}
	water := 0
	for i, h := range height {
		if leftMax[i] >= rightMax[i] && rightMax[i] > h {
			water += rightMax[i] - h
		} else if rightMax[i] > leftMax[i] && leftMax[i] > h {
			water += leftMax[i] - h
		}
	}
	return water
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	fmt.Printf("water is: %v", a)
}
