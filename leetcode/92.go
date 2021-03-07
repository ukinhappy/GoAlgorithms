package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if left == 1 {
		return reverseN(head, right)
	}
	head.Next = reverseBetween(head.Next, left-1, right-1)
	return head
}

var flag *ListNode

func reverseN(head *ListNode, n int) *ListNode {
	if n == 1 {
		flag = head.Next
		return head
	}
	newHead := reverseN(head.Next, n-1)
	head.Next.Next = head
	head.Next = flag
	return newHead
}
