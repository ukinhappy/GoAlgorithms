package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func isPalindrome(head *ListNode) bool {
	left = head
	return traverse(head)
}

var left *ListNode

func traverse(right *ListNode) bool {
	if right == nil {
		return true
	}

	result := traverse(right.Next)
	result = (left.Val == right.Val) && result
	left = left.Next
	return result
}
