package leetcode860

import "fmt"

func lemonadeChange(bills []int) bool {
	i := 0
	five := 0
	ten := 0
	twenty := 0
	for i < len(bills) {
		money := bills[i]
		if money == 5 {
			five++
			i++
		} else if money == 10 {
			if five > 0 {
				five--
				ten++
				i++
			} else {
				return false
			}
		} else {
			if ten > 0 && five > 0 {
				ten--
				five--
				twenty++
				i++
			} else if five >= 3 {
				five -= 3
				twenty++
				i++
			} else {
				return false
			}
		}
	}
	return true
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := lemonadeChange([]int{5, 5, 5, 10, 20})
	fmt.Printf("ans is: %v", ans)
}
