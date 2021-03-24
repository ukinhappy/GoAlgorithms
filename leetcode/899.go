package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructFromPrePost(pre []int, post []int) *TreeNode {
	//后序遍历的index
	postindex := make(map[int]int)
	for k, v := range post {
		postindex[v] = k
	}
	return tran(pre, post, 0, len(pre)-1, 0, len(post), postindex)
}
func tran(pre, post []int, prebegin, preend, postbegin, postend int, postindex map[int]int) *TreeNode {
	if prebegin > preend || postbegin > postend {
		return nil
	}
	if prebegin == preend{
		return &TreeNode{Val: pre[prebegin]}
	}
	tmpindex := postindex[pre[prebegin+1]]

	var root = &TreeNode{Val: pre[prebegin]}
	root.Left = tran(pre, post, prebegin+1, prebegin+tmpindex-postbegin+1, postbegin, tmpindex, postindex)
	root.Right = tran(pre, post, prebegin+tmpindex-postbegin+1+1, preend, tmpindex+1, postend-1, postindex)

	return root
}
