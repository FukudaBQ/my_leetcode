package leetcode203

import (
	"by_golang/listnode"
	"fmt"
)

func removeElements(head *listnode.ListNode, val int) *listnode.ListNode {
	pre := &listnode.ListNode{val + 1, head}
	keptVal := pre
	pt := head
	for pt != nil {
		if pt.Val == val {
			pre.Next = pt.Next
			pt = pt.Next
		} else {
			pre = pt
			pt = pt.Next
		}
	}
	return keptVal.Next
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := removeElements(listnode.BuildList([]int{1, 2, 6, 3, 4, 5, 6}), 6)
	listnode.PrintList(a)
}
