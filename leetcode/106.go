package leetcode

func buildTree(inorder []int, postorder []int) *TreeNode {
	iindex := make(map[int]int)
	for k, v := range inorder {
		iindex[v] = k
	}
	return build(inorder, postorder, 0, len(inorder)-1, 0, len(postorder)-1, iindex)
}

func build(inorder, postorder []int, ibegin, iend, pbegin, pend int, iindex map[int]int) *TreeNode {
	if ibegin > iend || pbegin > pend {
		return nil
	}
	//后序遍历最后一个元素 即根结点
	root := &TreeNode{Val: postorder[pend]}
	index := iindex[postorder[pend]]
	root.Left = build(inorder, postorder, ibegin, index-1, pbegin, pbegin+index-ibegin-1, iindex)
	root.Right = build(inorder, postorder, index+1, iend, pbegin+index-ibegin, pend-1, iindex)
	return root
}
