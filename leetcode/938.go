package leetcode

func rangeSumBST(root *TreeNode, low int, high int) int {

	if root == nil {
		return 0
	}

	var result int
	if root.Val >= low && root.Val <= high {
		result = root.Val
	}

	return result + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)

}
