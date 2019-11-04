package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	first, second := head, head

	for second != nil && second.Next != nil && second.Next.Next != nil {
		first = first.Next
		second = second.Next.Next
	}

	second = first.Next
	first.Next = nil

	return merge(sortList(head), sortList(second))
}

func merge(node1, node2 *ListNode) *ListNode {

	var head, tmp *ListNode
	if node1 == nil {
		return node2
	}
	if node2 == nil {
		return node1
	}

	if node1.Val > node2.Val {
		head = node2
		node2 = node2.Next
	} else {
		head = node1
		node1 = node1.Next
	}
	tmp = head

	for node1 != nil && node2 != nil {
		if node1.Val > node2.Val {
			tmp.Next = node2
			node2 = node2.Next
		} else {
			tmp.Next = node1
			node1 = node1.Next
		}
		tmp = tmp.Next
	}

	if node1 != nil {
		tmp.Next = node1
	}

	if node2 != nil {
		tmp.Next = node2
	}

	return head
}
