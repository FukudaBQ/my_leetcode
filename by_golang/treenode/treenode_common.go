package treenode

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree 按照 LeetCode 的层序数组构建二叉树。
// 使用 []any，nil 表示空节点，例如：[]any{3, 9, 20, nil, nil, 15, 7}
func BuildTree(nums []any) *TreeNode {
	if len(nums) == 0 || nums[0] == nil {
		return nil
	}
	root := &TreeNode{Val: toInt(nums[0])}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(nums) {
		node := queue[0]
		queue = queue[1:]

		if i < len(nums) && nums[i] != nil {
			node.Left = &TreeNode{Val: toInt(nums[i])}
			queue = append(queue, node.Left)
		}
		i++

		if i < len(nums) && nums[i] != nil {
			node.Right = &TreeNode{Val: toInt(nums[i])}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

// PrintTree 按层序打印二叉树，缺失节点用 null 占位（LeetCode 风格）。
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("[]")
		return
	}
	queue := []*TreeNode{root}
	parts := make([]string, 0)
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		if node == nil {
			parts = append(parts, "null")
			continue
		}
		parts = append(parts, fmt.Sprintf("%d", node.Val))
		queue = append(queue, node.Left, node.Right)
	}
	// 去掉末尾多余的 null
	for len(parts) > 0 && parts[len(parts)-1] == "null" {
		parts = parts[:len(parts)-1]
	}
	fmt.Println("[" + strings.Join(parts, ",") + "]")
}

func toInt(v any) int {
	switch x := v.(type) {
	case int:
		return x
	case int32:
		return int(x)
	case int64:
		return int(x)
	case float64:
		return int(x)
	default:
		panic(fmt.Sprintf("treenode: unsupported value type %T", v))
	}
}
