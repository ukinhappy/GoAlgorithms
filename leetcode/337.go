package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

var m map[*TreeNode]int = make(map[*TreeNode]int, 0)

func rob(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if m[root] > 0 {
		return m[root]
	}
	doit := root.Val
	if root.Left != nil {
		doit += rob(root.Left.Left) + rob(root.Left.Right)
	}
	if root.Right != nil {
		doit += rob(root.Right.Left) + rob(root.Right.Right)
	}

	dontdoit := rob(root.Left) + rob(root.Right)
	m[root] = max(doit, dontdoit)
	return m[root]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
