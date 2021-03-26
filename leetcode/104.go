package leetcode

//深度优先遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right))+1
}
func max(l, r int) int {
	if l > r {
		return l
	}
	return r
}

//广度优先遍历
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	list := []*TreeNode{root}

	var result = 0
	for len(list) > 0 {
		result++
		var tmplist []*TreeNode

		for len(list) > 0 {
			tmp := list[0]
			list = list[1:]
			if tmp.Left != nil {
				tmplist = append(tmplist, tmp.Left)
			}
			if tmp.Right != nil {
				tmplist = append(tmplist, tmp.Right)
			}
		}

		list = tmplist

	}
	return result
}
