package leetcode

import list2 "container/list"

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var flag *TreeNode
	list := list2.List{}
	for list.Len() > 0 || root != nil {
		if root != nil {
			list.PushBack(root)
			root = root.Left
		} else {
			back := list.Back()
			if back.Value.(*TreeNode).Right != nil && back.Value.(*TreeNode).Right != flag {
				root = back.Value.(*TreeNode).Right
			} else {
				list.Remove(back)
				result = append(result, back.Value.(*TreeNode).Val)
				flag = back.Value.(*TreeNode)
				root = nil

			}
		}
	}

	return result
}
