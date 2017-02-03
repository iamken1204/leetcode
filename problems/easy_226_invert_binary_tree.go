package problems

// TreeNode had been declared at easy_104

// InvertTree represents
// Easy 226. Invert Binary Tree
// https://leetcode.com/problems/invert-binary-tree/
func InvertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = InvertTree(root.Right), InvertTree(root.Left)
	return root
}
