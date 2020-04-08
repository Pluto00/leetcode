package main

/*
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-linked-list-elements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	var (
		dummy = &ListNode{0, nil}
		cur   = dummy
	)
	for p := head; p != nil; p = p.Next {
		if p.Val != val {
			cur.Next = p
			cur = cur.Next
		}
	}
	cur.Next = nil
	return dummy.Next
}
