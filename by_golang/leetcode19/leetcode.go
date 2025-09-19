package leetcode19

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

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// keptHead := head
	pt := head
	ptlag := head
	pre := &ListNode{0, head}
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

	a := removeNthFromEnd(buildList([]int{1, 2, 3, 4, 5}), 2)
	printList(a)
}
