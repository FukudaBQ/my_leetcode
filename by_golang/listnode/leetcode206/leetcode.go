package leetcode206

import (
	"by_golang/listnode"
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *listnode.ListNode) *listnode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var last *listnode.ListNode
	for head.Next != nil {
		rawHead := head
		head = head.Next
		rawHead.Next = last
		last = rawHead
	}
	head.Next = last
	return head
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := reverseList(listnode.BuildList([]int{1, 2, 3, 4, 5}))
	listnode.PrintList(a)
}
