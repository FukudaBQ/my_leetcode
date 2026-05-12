package leetcode222

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
func countNodes(root *treenode.TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth, rightDepth := 0, 0
	for node := root; node != nil; node = node.Left {
		leftDepth++
	}
	for node := root; node != nil; node = node.Right {
		rightDepth++
	}
	if leftDepth == rightDepth {
		return 1<<leftDepth - 1
	}
	return countNodes(root.Left) + countNodes(root.Right) + 1
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := countNodes(treenode.BuildTree([]any{1, 2, 3, 4, 5, 6}))
	fmt.Printf("The max depth of tree is: %v\n", a)
}
