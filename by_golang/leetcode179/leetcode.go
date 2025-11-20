package leetcode179

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	strNums := make([]string, len(nums))
	for i, n := range nums {
		strNums[i] = strconv.Itoa(n)
	}
	sort.Slice(strNums, func(i, j int) bool {
		a, b := strNums[i], strNums[j]
		return a+b > b+a
	})
	if strNums[0] == "0" {
		return "0"
	}
	res := strings.Builder{}
	res.Grow(len(nums))
	for _, n := range strNums {
		res.WriteString(n)
	}
	return res.String()
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := largestNumber([]int{3, 30, 34, 5, 9})
	fmt.Println(a)
}
