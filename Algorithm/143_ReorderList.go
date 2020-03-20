package main

/*
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list/
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}
	var (
		helper       func(*ListNode, int) *ListNode
		lengthOfList int
	)
	helper = func(node *ListNode, len int) *ListNode {
		if len == 1 {
			return node
		}
		if len == 2 {
			return node.Next
		}
		tail := helper(node.Next, len-2)
		tmp := tail.Next
		tail.Next = tail.Next.Next
		tmp.Next = node.Next
		node.Next = tmp
		return tail
	}
	for cur := head; cur != nil; cur = cur.Next {
		lengthOfList++
	}
	helper(head, lengthOfList)
}
