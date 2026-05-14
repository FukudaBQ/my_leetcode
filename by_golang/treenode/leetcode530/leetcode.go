package leetcode530

import (
	"by_golang/treenode"
	"fmt"
	"math"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *treenode.TreeNode) int {
	low := math.MaxInt
	var prev *treenode.TreeNode
	var inorder func(node *treenode.TreeNode)
	inorder = func(node *treenode.TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		if prev != nil && node.Val-prev.Val < low {
			low = node.Val - prev.Val
		}
		prev = node
		inorder(node.Right)
	}
	inorder(root)
	return low
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	//a := getMinimumDifference(treenode.BuildTree([]any{4, 2, 6, 1, 3}))
	//a := getMinimumDifference(treenode.BuildTree([]any{1, 0, 48, nil, nil, 12, 49}))
	a := getMinimumDifference(treenode.BuildTree([]any{1, nil, 3, 2}))
	fmt.Printf("The max depth of tree is: %v\n", a)

}
