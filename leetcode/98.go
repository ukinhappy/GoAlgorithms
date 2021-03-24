package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return valid(root, nil, nil)
}

func valid(root, min, max *TreeNode) bool {
	if root == nil {
		return true
	}
	if max != nil && root.Val > max.Val {
		return false
	}
	if min != nil && root.Val < min.Val {
		return false
	}
	return valid(root.Left, min, root) && valid(root.Right, root, max)
}
