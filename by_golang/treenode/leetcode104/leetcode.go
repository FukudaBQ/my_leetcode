package leetcode104

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
func maxDepth(root *treenode.TreeNode) int {
	maxLevel := 0
	if root == nil {
		return 0
	}
	queue := []*treenode.TreeNode{root}
	levelQueue := []int{0}
	for len(queue) > 0 {
		current := queue[0]
		currentLevel := levelQueue[0]
		queue = queue[1:]
		levelQueue = levelQueue[1:]
		if currentLevel > maxLevel {
			maxLevel = currentLevel
		}
		if current.Left != nil {
			queue = append(queue, current.Left)
			levelQueue = append(levelQueue, currentLevel+1)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
			levelQueue = append(levelQueue, currentLevel+1)
		}
	}
	return maxLevel + 1
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	a := maxDepth(treenode.BuildTree([]any{3, 9, 20, nil, nil, 15, 7}))
	fmt.Printf("The max depth of tree is: %d\n", a)
}
