package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	n1, n2 := head.Next, head.Next.Next
	for {
		if n1 == nil || n2 == nil || n2.Next == nil {
			break
		}
		if n1 == n2 {
			return true
		}
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	return false
}
