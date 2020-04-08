package main

/*
给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。

本题中，一个高度平衡二叉树是指一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/convert-sorted-list-to-binary-search-tree
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{
			Val:   head.Val,
			Left:  nil,
			Right: nil,
		}
	}
	midPoint := findMidPointFromList(head)
	return &TreeNode{
		Val:   midPoint.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(midPoint.Next),
	}
}

func findMidPointFromList (head *ListNode) *ListNode {
	var fastPoint, slowPoint, prePoint = head, head, head
	for fastPoint != nil && fastPoint.Next != nil {
		prePoint = slowPoint
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
	}
	prePoint.Next = nil
	return slowPoint
}