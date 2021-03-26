package leetcode

type Node struct {
	Val      int
	Children []*Node
}

//深度遍历
func postorder(root *Node) []int {
	if root == nil {
		return nil
	}

	var result []int
	for _, v := range root.Children {
		result = append(result, postorder(v)...)
	}
	result = append(result, root.Val)

	return result
}
