package leetcode605

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	k := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			back := true
			if i-1 >= 0 {
				back = flowerbed[i-1] == 0
			}
			front := true
			if i+1 < len(flowerbed) {
				front = flowerbed[i+1] == 0
			}
			if back && front {
				k++
				flowerbed[i] = 1
			}
		}
	}
	return k >= n
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)
	ans := canPlaceFlowers([]int{1, 0, 0, 0, 0, 1}, 2)
	fmt.Printf("ans is: %v", ans)
}
