package leetcode19

import (
	"by_golang/listnode"
	"fmt"
)

func removeNthFromEnd(head *listnode.ListNode, n int) *listnode.ListNode {
	// keptHead := head
	pt := head
	ptlag := head
	pre := &listnode.ListNode{0, head}
	for ; n >= 0; n-- {
		pt = pt.Next
	}
	for pt != nil {
		pt = pt.Next
		ptlag = ptlag.Next
		pre = pre.Next
	}

	return ptlag
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := removeNthFromEnd(listnode.BuildList([]int{1, 2, 3, 4, 5}), 2)
	listnode.PrintList(a)
}
