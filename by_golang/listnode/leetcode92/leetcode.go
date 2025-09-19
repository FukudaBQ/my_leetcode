package leetcode92

import (
	"by_golang/listnode"
	"fmt"
)

func reverseBetween(head *listnode.ListNode, left int, right int) *listnode.ListNode {
	pre := &listnode.ListNode{1, head}
	ptL := pre
	ptR := pre
	for i := 0; i < left-1; i++ {
		ptL = ptL.Next
		ptR = ptR.Next
	}
	for i := 0; i < right-left+1; i++ {
		ptR = ptR.Next
	}

	return ptR
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := reverseBetween(listnode.BuildList([]int{1, 2, 3, 4, 5, 6}), 1, 2)
	listnode.PrintList(a)
}
