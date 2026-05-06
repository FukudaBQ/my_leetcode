package leetcode103

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
func zigzagLevelOrder(root *treenode.TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	queue := []*treenode.TreeNode{root}
	levelQueue := []int{0}
	for len(queue) > 0 {
		current := queue[0]
		currentLevel := levelQueue[0]
		queue = queue[1:]
		levelQueue = levelQueue[1:]
		if len(res) <= currentLevel {
			res = append(res, make([]int, 0))
		}
		res[currentLevel] = append(res[currentLevel], current.Val)
		if current.Left != nil {
			queue = append(queue, current.Left)
			levelQueue = append(levelQueue, currentLevel+1)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
			levelQueue = append(levelQueue, currentLevel+1)
		}
	}
	for i, j := range res {
		if i%2 == 1 {
			slices.Reverse(j)
		}
	}
	return res
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	root := treenode.BuildTree([]any{3, 9, 20, nil, nil, 15, 7})
	treenode.PrintTree(root)

	a := zigzagLevelOrder(root)
	fmt.Printf("levelOrder is: %v\n", a)
}
