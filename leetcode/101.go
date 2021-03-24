package leetcode

//递归方式
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return check(root.Left, root.Right)
}

func check(left, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return check(left.Left, right.Right) && check(left.Right, right.Left)
}

// 层级遍历判断
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	list := []*TreeNode{root}

	for {
		begin := 0
		end := len(list) - 1
		for begin >= end {
			if list[begin].Val != list[end].Val {
				return false
			}
			if list[begin].Left != nil && list[end].Right == nil {
				return false
			}
			if list[begin].Left == nil && list[end].Right != nil {
				return false
			}

			if list[begin].Right != nil && list[end].Left == nil {
				return false
			}
			if list[begin].Right == nil && list[end].Left != nil {
				return false
			}
			begin++
			end--
		}

		// 暂存下一层结点
		var nextTmplist []*TreeNode
		for i := 0; i < len(list); i++ {
			if list[i].Left != nil {
				nextTmplist = append(nextTmplist, list[i].Left)
			}
			if list[i].Right != nil {
				nextTmplist = append(nextTmplist, list[i].Right)
			}
		}

		list = nextTmplist

		if len(nextTmplist) == 0 {
			return true
		}
	}
}

// 优化后的层级遍历判断
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	list := []*TreeNode{root.Left, root.Right}
	for len(list) > 0 {
		begin := list[0]
		end := list[1]

		list = list[2:]
		if begin == nil && end == nil {
			continue
		}
		if begin == nil || end == nil {
			return false
		}

		if begin.Val != end.Val {
			return false
		}
		list = append(list, begin.Left, end.Right, begin.Right, end.Left)
	}
	return true
}
