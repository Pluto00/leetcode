package main

/*
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	var (
		dummy1 = &ListNode{0, nil}
		dummy2 = &ListNode{0, nil}
		cur1   = dummy1
		cur2   = dummy2
	)
	for p := head; p != nil; p = p.Next {
		if p.Val < x {
			cur1.Next = p
			cur1 = p
		} else {
			cur2.Next = p
			cur2 = p
		}
	}
	cur1.Next = dummy2.Next
	cur2.Next = nil
	return dummy1.Next
}
