package leetcode236

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
func lowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	return right
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := lowestCommonAncestor(treenode.BuildTree([]any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4}), treenode.BuildTree([]any{5, 6, 2, nil, nil, 7, 4}), treenode.BuildTree([]any{4}))
	fmt.Printf("The max depth of tree is: %v\n", a)
}
