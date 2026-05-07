package leetcode226

import (
	"by_golang/treenode"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *treenode.TreeNode) *treenode.TreeNode {
	if root == nil {
		return nil
	}
	stack := []*treenode.TreeNode{root}
	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]
		node.Left, node.Right = node.Right, node.Left
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}

	}
	return root
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	root := treenode.BuildTree([]any{4, 2, 7, 1, 3, 6, 9})
	treenode.PrintTree(root)

	a := invertTree(root)
	treenode.PrintTree(a)
}
