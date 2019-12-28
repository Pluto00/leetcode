package main

/*
给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list
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
	if head == nil {
		return head
	}
	pi := head
	for i := head.Next; i != nil; i = i.Next {
		if pi.Val == i.Val {
			pi.Next = i.Next
		} else {
			pi = pi.Next
		}
	}
	return head
}
