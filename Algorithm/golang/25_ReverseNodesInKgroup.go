package main

/*
给你一个链表，每 k 个节点一组进行翻转，请你返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。

如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-nodes-in-k-group
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		pre, nextHead,
		next, cur *ListNode
		dummy = &ListNode{0, nil}
	)
	pre = dummy
	cur = head
	for cur != nil {
		tmp, i := cur, 0
		for i = 0; i < k && tmp != nil; i++ {
			tmp = tmp.Next
		}
		if i < k {
			pre.Next = cur
			break
		}
		nextHead = cur
		for i = 0; i < k && cur != nil; i++ {
			next = cur.Next
			cur.Next = pre.Next
			pre.Next = cur
			cur = next
		}
		pre = nextHead
	}
	return dummy.Next
}
