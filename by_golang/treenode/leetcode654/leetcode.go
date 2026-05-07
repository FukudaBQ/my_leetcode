package leetcode654

import (
	"by_golang/treenode"
	"fmt"
	"slices"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *treenode.TreeNode {
	if nums == nil || len(nums) == 0 {
		return nil
	}
	current := slices.Max(nums)
	index := slices.Index(nums, current)
	left := nums[:index]
	right := nums[index+1:]
	return &treenode.TreeNode{
		Val:   current,
		Left:  constructMaximumBinaryTree(left),
		Right: constructMaximumBinaryTree(right),
	}
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := constructMaximumBinaryTree([]int{3, 2, 1, 6, 0, 5})
	treenode.PrintTree(a)
}
