package leetcode538

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
func convertBST(root *treenode.TreeNode) *treenode.TreeNode {
	res, _ := greaterTree(root, 0)
	return res
}

func greaterTree(root *treenode.TreeNode, input int) (tree *treenode.TreeNode, leftest int) {
	if root == nil {
		return nil, 0
	}
	var right *treenode.TreeNode
	var left *treenode.TreeNode
	leftest = input
	if root.Right != nil {
		right, leftest = greaterTree(root.Right, input)
		root.Right = right
	}
	root.Val = root.Val + leftest
	leftest = root.Val
	if root.Left != nil {
		left, leftest = greaterTree(root.Left, root.Val)
		root.Left = left
	}
	return root, leftest
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := convertBST(treenode.BuildTree([]any{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}))
	treenode.PrintTree(a)

}
