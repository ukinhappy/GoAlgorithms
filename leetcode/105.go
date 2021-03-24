package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {

	iindex := make(map[int]int)

	for k, v := range inorder {
		iindex[v] = k
	}
	return tran(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1, iindex)
}


func tran(preorder []int, inorder []int, pbegin, pend, ibegin, iend int, iindex map[int]int) *TreeNode {
	if pbegin > pend || ibegin > iend {
		return nil
	}
	var rootvalue = preorder[pbegin]
	var root = &TreeNode{Val: rootvalue}
	root.Left = tran(preorder, inorder, pbegin+1, pbegin+iindex[rootvalue]-ibegin, ibegin, iindex[rootvalue]-1, iindex)
	root.Right = tran(preorder, inorder, pbegin+iindex[rootvalue]-ibegin+1, pend, iindex[rootvalue]+1, iend, iindex)
	return root
}

