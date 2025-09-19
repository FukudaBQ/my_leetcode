package leetcode160

import (
	"by_golang/listnode"
	"fmt"
)

func getIntersectionNode(headA, headB *listnode.ListNode) *listnode.ListNode {
	ptr1 := headA
	ptr2 := headB
	for ptr1 != ptr2 {
		if ptr1 != nil {
			ptr1 = ptr1.Next
		} else {
			ptr1 = headB
		}
		if ptr2 != nil {
			ptr2 = ptr2.Next
		} else {
			ptr2 = headA
		}
	}
	return ptr1
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := getIntersectionNode(listnode.BuildList([]int{4, 1, 8, 4, 5}), listnode.BuildList([]int{5, 6, 1, 8, 4, 5}))
	listnode.PrintList(a)
}
