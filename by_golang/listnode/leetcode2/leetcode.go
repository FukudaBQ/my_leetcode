package leetcode2

import (
	"by_golang/listnode"
	"fmt"
)

func addTwoNumbers(l1 *listnode.ListNode, l2 *listnode.ListNode) *listnode.ListNode {
	res := &listnode.ListNode{1, nil}
	keptRes := res
	pt1 := l1
	pt2 := l2
	taken := false
	for pt1 != nil && pt2 != nil {
		current := pt1.Val + pt2.Val
		if taken {
			current++
		}
		taken = current >= 10
		res.Next = &listnode.ListNode{current % 10, nil}
		res = res.Next
		pt1 = pt1.Next
		pt2 = pt2.Next
	}
	var last *listnode.ListNode
	if pt1 != nil {
		last = pt1
	} else {
		last = pt2
	}
	for last != nil {
		current := last.Val
		if taken {
			current++
		}
		taken = current >= 10
		res.Next = &listnode.ListNode{current % 10, nil}
		res = res.Next
		last = last.Next
	}
	if taken {
		res.Next = &listnode.ListNode{1, nil}
	}
	return keptRes.Next
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := addTwoNumbers(listnode.BuildList([]int{2, 4, 3}), listnode.BuildList([]int{5, 6, 4}))
	listnode.PrintList(a)
}
