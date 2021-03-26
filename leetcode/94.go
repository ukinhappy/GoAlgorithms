package leetcode

import (
	list2 "container/list"
)

var result []int

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return result
	}
	inorderTraversal(root.Left)
	result = append(result, root.Val)
	inorderTraversal(root.Right)
	return result

}

//广度优先遍历
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	list := list2.List{}
	for root != nil || list.Len() > 0 {
		if root != nil {
			list.PushBack(root)
			root = root.Left
		} else {
			back := list.Back()
			list.Remove(back)
			result = append(result, back.Value.(*TreeNode).Val)
			root = back.Value.(*TreeNode).Right
		}
	}
	return result

}
