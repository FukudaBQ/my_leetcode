package leetcode135

import "fmt"

func candy(ratings []int) int {
	candies := make([]int, len(ratings))
	candies[0] = 1
	for i := 1; i < len(candies); i++ {
		if ratings[i] > ratings[i-1] {
			candies[i] = candies[i-1] + 1
		} else {
			candies[i] = 1
		}
	}
	for i := len(candies) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			if candies[i+1]+1 > candies[i] {
				candies[i] = candies[i+1] + 1
			}
		}
	}
	sum := 0
	for _, v := range candies {
		sum += v
	}
	return sum
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := candy([]int{2, 3, 1, 4, 5, 6, 0, 7, 2, 1, 3, 4, 5})
	fmt.Printf("is anagram: %v", a)
}
