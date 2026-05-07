package leetcode105

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
func buildTree(preorder []int, inorder []int) *treenode.TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := preorder[0]
	left := slices.Index(inorder, root)
	if left == -1 {
		return nil
	}
	leftInorder := inorder[:left]
	rightInorder := inorder[left+1:]
	leftPreorder := preorder[1 : left+1]
	rightPreorder := preorder[left+1:]
	return &treenode.TreeNode{
		Val:   root,
		Left:  buildTree(leftPreorder, leftInorder),
		Right: buildTree(rightPreorder, rightInorder),
	}
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	root := treenode.BuildTree([]any{3, 9, 20, nil, nil, 15, 7})
	treenode.PrintTree(root)

	a := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})
	treenode.PrintTree(a)
}
