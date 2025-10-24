package leetcode167

import "fmt"

func twoSum(numbers []int, target int) []int {
	pt1 := 0
	pt2 := len(numbers) - 1
	for pt2 > pt1 {
		sum := numbers[pt1] + numbers[pt2]
		if sum == target {
			return []int{pt1 + 1, pt2 + 1}
		}
		if sum > target {
			pt2--
		} else {
			pt1++
		}
	}
	return nil
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Printf("ans is: %v", ans)
}
