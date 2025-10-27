package leetcode142

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
func detectCycle(head *listnode.ListNode) *listnode.ListNode {
	slow, fast := head, head
	for fast != nil && slow != nil {
		if fast.Next != nil && fast.Next.Next != nil {
			fast = fast.Next.Next
			slow = slow.Next
		} else {
			return nil
		}
		if fast == slow {
			break
		}
	}
	slow = head
	for fast != nil && slow != nil {
		if slow == fast {
			return slow
		}
		slow = slow.Next
		fast = fast.Next
	}
	return nil
}

func Leetcode() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	_ = detectCycle(BuildCycleList([]int{3, 2, 0, -4}, 1))
	//listnode.PrintList(a)
}

func BuildCycleList(nums []int, pos int) *listnode.ListNode {
	head := listnode.BuildList(nums)
	if head == nil || pos < 0 || pos >= len(nums) {
		return head
	}

	tail := head
	for tail.Next != nil {
		tail = tail.Next
	}

	entry := head
	for i := 0; i < pos; i++ {
		entry = entry.Next
	}

	tail.Next = entry
	return head
}
