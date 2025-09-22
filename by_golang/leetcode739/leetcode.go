package leetcode739

import "fmt"

func dailyTemperatures(temperatures []int) []int {
	res := make([]int, len(temperatures))
	stack := make([]int, 0)
	for i, temp := range temperatures {
		if len(stack) == 0 || temperatures[stack[len(stack)-1]] >= temp {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && temperatures[stack[len(stack)-1]] < temp {
			res[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := dailyTemperatures([]int{34, 80, 80, 34, 34, 80, 80, 80, 80, 34})
	fmt.Printf("temperature is: %v", a)
}
