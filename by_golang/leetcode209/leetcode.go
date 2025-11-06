package leetcode209

import "fmt"

func minSubArrayLen(target int, nums []int) int {
	pt0, pt1 := 0, 1
	sum := 0
	res := len(nums)
	sum += nums[0]
	for pt1 < len(nums) {
		//fmt.Printf("pt0 pt1 sum: %d %d %d\n", pt0, pt1, sum)
		if sum < target || sum == 0 {
			sum += nums[pt1]
			pt1++
		} else {
			length := pt1 - pt0
			//fmt.Printf("pt1 length: %d\n", length)
			if length < res {
				res = length
			}
			sum -= nums[pt0]
			pt0++
		}
	}
	if res == len(nums) && sum < target {
		return 0
	}
	for sum >= target {
		if pt1-pt0 < res {
			res = pt1 - pt0
		}
		sum -= nums[pt0]
		pt0++
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	a := minSubArrayLen(15, []int{1, 2, 3, 4, 5})
	fmt.Printf("is anagram: %v", a)
}
