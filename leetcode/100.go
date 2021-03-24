package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//深度优先遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

//广度优先遍历
func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil {
		return false
	}
	if q == nil {
		return false
	}

	queue1 := []*TreeNode{p}
	queue2 := []*TreeNode{q}
	for len(queue2) > 0 && len(queue1) > 0 {
		ptmp := queue1[0]
		qtmp := queue2[0]
		queue1 = queue1[1:]
		queue2 = queue2[1:]
		if ptmp.Val != qtmp.Val {
			return false
		}

		// 判断左右是否都有
		if (ptmp.Left != nil && qtmp.Left == nil) ||
			(ptmp.Left == nil && qtmp.Left != nil) ||
			(ptmp.Right != nil && qtmp.Right == nil) ||
			(ptmp.Right == nil && qtmp.Right != nil) {
			return false
		}

		//都有的话，把左右直接分别投递到队列中，等待取出判断是否相等
		if ptmp.Left != nil {
			queue1 = append(queue1, ptmp.Left)
		}
		if ptmp.Right != nil {
			queue1 = append(queue1, ptmp.Right)
		}
		if qtmp.Left != nil {
			queue2 = append(queue2, qtmp.Left)
		}
		if qtmp.Right != nil {
			queue2 = append(queue2, qtmp.Right)
		}

	}
	return len(queue1) == len(queue2)
}
