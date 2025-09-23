package leetcode84

import "fmt"

func largestRectangleArea(heights []int) int {
	length := make([]int, len(heights))
	stack := make([]int, 0)
	for i, height := range heights {
		if len(stack) == 0 || height <= heights[stack[len(stack)-1]] {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && heights[stack[len(stack)-1]] <= height {
			length[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}

	lengthback := make([]int, len(heights))
	stackback := make([]int, 0)
	for i, height := range heights {
		if len(stackback) == 0 || height >= heights[stackback[len(stackback)-1]] {
			stackback = append(stackback, i)
			continue
		}
		for len(stackback) > 0 && heights[stackback[len(stackback)-1]] >= height {
			lengthback[stackback[len(stackback)-1]] = i - stackback[len(stackback)-1]
			stackback = stackback[:len(stackback)-1]
		}
		stackback = append(stackback, i)
	}
	area := 0
	for i, height := range heights {
		if (lengthback[i]+length[i]-1)*height > area {
			area = (lengthback[i] + length[i] - 1) * height
		}
	}
	return area
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	fmt.Printf("height is: %v", a)
}
