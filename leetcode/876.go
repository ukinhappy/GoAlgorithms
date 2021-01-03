package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {

	n1, n2 := head, head
	for {
		if n1 == nil || n2 == nil || n2.Next == nil {
			break
		}
		n1 = n1.Next
		n2 = n2.Next.Next
	}

	return n1
}
