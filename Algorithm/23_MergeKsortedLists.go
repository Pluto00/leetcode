package main

/*
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	amount := len(lists)
	interval := 1
	for interval < amount {
		for i := 0; i < amount-interval; i += 2 * interval {
			lists[i] = merge2List(lists[i], lists[i+interval])
		}
		interval *= 2
	}
	return lists[0]
}

func merge2List(l1, l2 *ListNode) *ListNode {
	head := &ListNode{0, nil}
	point := head
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			point.Next = l1
			l1 = l1.Next
		} else {
			point.Next = l2
			l2 = l2.Next
		}
		point = point.Next
	}
	if l1 != nil {
		point.Next = l1
	} else {
		point.Next = l2
	}
	return head.Next
}
