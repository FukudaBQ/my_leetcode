package leetcode199

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
func rightSideView(root *treenode.TreeNode) []int {
	if root == nil {
		return nil
	}
	result := make([]int, 0)
	queue := []*treenode.TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		rightmost := queue[0].Val
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			rightmost = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, rightmost)
	}
	return result
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := rightSideView(treenode.BuildTree([]any{1, 2, 3, 4, nil, nil, nil, 5}))
	fmt.Printf("The max depth of tree is: %v\n", a)
}
