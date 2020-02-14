package main

/*
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var (
		dummy = &ListNode{head.Val - 1, head}
		pre   = dummy
	)
	for p1 := head; p1 != nil; p1 = p1.Next {
		if p1.Next != nil && p1.Next.Val == p1.Val {
			for p1 = p1.Next; p1.Next != nil && p1.Next.Val == p1.Val; p1 = p1.Next {
			}
		} else {
			pre.Next = p1
			pre = pre.Next
		}
	}
	pre.Next = nil
	return dummy.Next
}
