package leetcode98

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
func isValidBST(root *treenode.TreeNode) bool {
	return validateBST(root, nil, nil)
}

func validateBST(root *treenode.TreeNode, min, max *int) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Val >= *max {
		return false
	}
	if min != nil && root.Val <= *min {
		return false
	}
	return validateBST(root.Left, min, &root.Val) && validateBST(root.Right, &root.Val, max)
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := isValidBST(treenode.BuildTree([]any{5, 4, 6, nil, nil, 3, 7}))
	fmt.Printf("The max depth of tree is: %v\n", a)
}
