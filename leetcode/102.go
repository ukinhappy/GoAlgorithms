package leetcode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	list := []*TreeNode{root}
	for {
		var tmplist []*TreeNode
		var tmpvalue []int
		for len(list) > 0 {
			tmp := list[0]
			tmpvalue = append(tmpvalue, tmp.Val)
			if tmp.Left != nil {
				tmplist = append(tmplist, tmp.Left)
			}
			if tmp.Right != nil {
				tmplist = append(tmplist, tmp.Right)
			}
			list = list[1:]
		}
		result = append(result, tmpvalue)
		list = tmplist

		if len(list) == 0 {
			return result
		}
	}
}
