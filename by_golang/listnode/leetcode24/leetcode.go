package leetcode24

import (
	"by_golang/listnode"
	"fmt"
)

func swapPairs(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return head
	}
	headNext := head.Next
	pre := &listnode.ListNode{0, head}
	keptHead := pre

	for headNext != nil && head != nil {
		head.Next = headNext.Next
		headNext.Next = head
		pre.Next = headNext

		pre = head
		head = head.Next
		if head != nil && head.Next != nil {
			headNext = head.Next
		} else {
			headNext = nil
		}
	}
	return keptHead.Next
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := swapPairs(listnode.BuildList([]int{1, 2, 3, 4}))
	listnode.PrintList(a)
}
