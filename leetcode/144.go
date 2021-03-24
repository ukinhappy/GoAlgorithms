package leetcode

import list2 "container/list"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var result []int

func preorderTraversal(root *TreeNode) []int {
	result = make([]int, 0)
	tran(root)
	return result
}

func tran(root *TreeNode) {
	if root == nil {
		return
	}

	result = append(result, root.Val)
	tran(root.Left)
	tran(root.Right)
}

// 非递归实现
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	result = make([]int, 0)
	list := list2.List{}

	for root != nil || list.Len() > 0 {

		if root != nil {
			result = append(result, root.Val)
			list.PushBack(root)
			root = root.Left
		} else {
			back := list.Back()
			list.Remove(back)
			root = back.Value.(*TreeNode).Right
		}
	}
	return result
}
