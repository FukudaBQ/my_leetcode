package leetcode21

import (
	"by_golang/listnode"
	"fmt"
)

func mergeTwoLists(list1 *listnode.ListNode, list2 *listnode.ListNode) *listnode.ListNode {
	pre := &listnode.ListNode{0, nil}
	head := pre
	ptr1 := list1
	ptr2 := list2
	for ptr1 != nil && ptr2 != nil {
		if ptr1.Val < ptr2.Val {
			pre.Next = ptr1
			pre = pre.Next
			ptr1 = ptr1.Next
		} else {
			pre.Next = ptr2
			pre = pre.Next
			ptr2 = ptr2.Next
		}
	}
	if ptr1 != nil {
		pre.Next = ptr1
	} else {
		pre.Next = ptr2
	}
	return head.Next
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := mergeTwoLists(listnode.BuildList([]int{1, 2, 4}), listnode.BuildList([]int{1, 3, 4}))
	listnode.PrintList(a)
}
