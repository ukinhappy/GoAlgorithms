package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	n1, n2 := head.Next, head.Next.Next
	var mid *ListNode
	for {
		if n1 == nil || n2 == nil || n2.Next == nil {
			break
		}
		if n1 == n2 {
			mid = n2
			break
		}
		n1 = n1.Next
		n2 = n2.Next.Next
	}
	if mid == nil {
		return nil
	}
	n1 = head
	for {
		if n1 == mid {
			return n1
		}
		n1 = n1.Next
		mid = mid.Next
	}
}
