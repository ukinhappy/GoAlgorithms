package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	headPre := &ListNode{Next: head}

	next := head
	for next != nil && next.Next != nil {
		tmp := next.Next
		if next.Val < tmp.Val {
			next = next.Next
			continue
		}
		// 找到需要往前插入的那个数据
		next.Next = tmp.Next
		hp, ha := headPre, headPre.Next
		for ha.Val < tmp.Val {
			hp = ha
			ha = ha.Next
		}
		hp.Next = tmp
		tmp.Next = ha

	}
	return headPre.Next
}
