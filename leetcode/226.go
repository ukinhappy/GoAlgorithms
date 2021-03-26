package leetcode

// 深度优先遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 广度优先遍历
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	list := []*TreeNode{root}
	for len(list) > 0 {
		tmp := list[0]
		list = list[1:]

		tmp.Left, tmp.Right = tmp.Right, tmp.Left
		if tmp.Left != nil {
			list = append(list, tmp.Left)
		}
		if tmp.Right != nil {
			list = append(list, tmp.Right)
		}
	}

	return root
}
