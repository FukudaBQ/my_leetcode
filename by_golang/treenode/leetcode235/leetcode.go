package leetcode235

import (
	"by_golang/treenode"
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */
func lowestCommonAncestor(root, p, q *treenode.TreeNode) *treenode.TreeNode {
	if p.Val < root.Val && q.Val < root.Val {
		return lowestCommonAncestor(root.Left, p, q)
	}
	if p.Val > root.Val && q.Val > root.Val {
		return lowestCommonAncestor(root.Right, p, q)
	}
	return root
}

func Leetcode() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	//a := getMinimumDifference(treenode.BuildTree([]any{4, 2, 6, 1, 3}))
	//a := getMinimumDifference(treenode.BuildTree([]any{1, 0, 48, nil, nil, 12, 49}))
	a := lowestCommonAncestor(treenode.BuildTree([]any{6, 2, 8, 0, 4, 7, 9, nil, nil, 3, 5}), treenode.BuildTree([]any{2}), treenode.BuildTree([]any{8}))
	treenode.PrintTree(a)

}
