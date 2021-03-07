package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	//
	newHead := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return newHead

}
// 遍历


func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var cur *ListNode
	for head!=nil{
		tmp:=head.Next
		head.Next=cur
		cur=head
		head=tmp
	}

	return cur

}
