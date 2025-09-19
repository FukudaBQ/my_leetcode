package leetcode24

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Printf("%d ", head.Val)
		head = head.Next
	}
	fmt.Println()
}

func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	headNext := head.Next
	pre := &ListNode{0, head}
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

	a := swapPairs(buildList([]int{1, 2, 3, 4}))
	printList(a)
}
