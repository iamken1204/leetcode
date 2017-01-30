package problems

import "math"

// TreeNode is the definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// MaxDepth represents
// 104. Maximum Depth of Binary Tree
// https://leetcode.com/problems/maximum-depth-of-binary-tree/
func MaxDepth(root *TreeNode) int {
	// The end of a tree is nil node.
	if root == nil {
		return 0
	}
	l := float64(MaxDepth(root.Left))
	r := float64(MaxDepth(root.Right))
	return 1 + int(math.Max(l, r))
}
