package leetcode1475

import (
	"fmt"
)

func finalPrices(prices []int) []int {
	compare := []int{0}
	var res []int
	for i := len(prices) - 1; i >= 0; i-- {
		price := prices[i]
		found := false
		for j := 0; j < len(compare); j++ {
			if price >= compare[j] {
				res = append([]int{price - compare[j]}, res...)
				compare = append([]int{price}, compare...)
				found = true
				break
			}
		}
		if !found {
			res = append([]int{price}, res...)
			compare = append([]int{price}, compare...)
		}
	}
	return res
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := finalPrices([]int{8, 7, 4, 2, 8, 1, 7, 7, 10, 1})
	fmt.Printf("Prices is: %v", a)
}
