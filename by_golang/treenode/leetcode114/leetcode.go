package leetcode114

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
func flatten(root *treenode.TreeNode) {
	if root == nil {
		return
	}
	preorders := make([]*treenode.TreeNode, 0)
	queue := []*treenode.TreeNode{root}
	for len(queue) > 0 {
		node := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		preorders = append(preorders, node)
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
	}
	if len(preorders) == 1 {
		return
	}
	extraFront := &treenode.TreeNode{}
	for i := 0; i < len(preorders); i++ {
		extraFront.Left = nil
		extraFront.Right = preorders[i]
		extraFront = extraFront.Right
	}
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	//a := treenode.BuildTree([]any{1, 2, 5, 3, 4, nil, 6})

	a := treenode.BuildTree([]any{0})

	flatten(a)
	treenode.PrintTree(a)
}
