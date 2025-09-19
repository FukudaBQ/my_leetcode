package leetcode328

import (
	"by_golang/listnode"
	"fmt"
)

func oddEvenList(head *listnode.ListNode) *listnode.ListNode {
	if head == nil {
		return head
	}
	pre := &listnode.ListNode{0, head}
	oddpt := head
	evenpt := head.Next
	keptEven := evenpt
	for oddpt != nil && evenpt != nil {
		oddpt.Next = evenpt.Next
		if oddpt.Next != nil {
			evenpt.Next = oddpt.Next.Next
		}
		if oddpt.Next == nil {
			oddpt.Next = keptEven
			oddpt = nil
			continue
		}
		if evenpt.Next == nil {
			oddpt = oddpt.Next
			oddpt.Next = keptEven
			oddpt = nil
			continue
		}
		oddpt = oddpt.Next
		evenpt = evenpt.Next
	}
	return pre.Next
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := oddEvenList(listnode.BuildList([]int{1, 2, 3, 4, 5}))
	listnode.PrintList(a)
}
