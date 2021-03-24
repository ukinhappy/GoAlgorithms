package leetcode

func levelOrderBottom(root *TreeNode) [][]int {

	var result [][]int

	tmp := levelOrder(root)
	if len(tmp) == 0 {
		return nil
	}

	for i := len(tmp) - 1; i >= 0; i-- {
		result = append(result,tmp[i])
	}
	return result
}
