package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	queue := []*TreeNode{root}

	for {
		if len(queue)<=0{
			break
		}

		tmp := make([]*TreeNode, 0)
		for _, n := range queue {
			if n.Left == nil && n.Right == nil {
				return depth
			}
			if n.Left != nil {
				tmp = append(tmp, n.Left)
			}
			if n.Right != nil {
				tmp = append(tmp, n.Right)
			}
		}

		queue = tmp
		depth++
	}

	return depth
}
