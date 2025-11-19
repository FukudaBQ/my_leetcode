package leetcode148

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
func sortList(head *listnode.ListNode) *listnode.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	fast, slow := head.Next, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	slowHead := slow.Next
	slow.Next = nil
	node1, node2 := sortList(head), sortList(slowHead)

	dummy := &listnode.ListNode{}
	pt := dummy
	for node1 != nil && node2 != nil {
		if node1.Val < node2.Val {
			pt.Next = node1
			node1 = node1.Next
		} else {
			pt.Next = node2
			node2 = node2.Next
		}
		pt = pt.Next
	}
	if node1 != nil {
		pt.Next = node1
	} else {
		pt.Next = node2
	}

	return dummy.Next
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := sortList(listnode.BuildList([]int{4, 2, 1, 3, 8, 5, 2, 6}))
	listnode.PrintList(a)
}
