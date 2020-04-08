package main

/*
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		length  = 0
		oldTail = head
		newTail = head
		newHead *ListNode
	)
	for length = 1; oldTail.Next != nil; length++ {
		oldTail = oldTail.Next
	}
	oldTail.Next = head

	k = k % length
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}
	newHead = newTail.Next
	newTail.Next = nil
	return newHead
}
