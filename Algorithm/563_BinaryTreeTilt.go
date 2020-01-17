package main

/*
给定一个二叉树，计算整个树的坡度。

一个树的节点的坡度定义即为，该节点左子树的结点之和和右子树结点之和的差的绝对值。空结点的的坡度是0。

整个树的坡度就是其所有节点的坡度之和。

注意:

任何子树的结点的和不会超过32位整数的范围。
坡度的值不会超过32位整数的范围。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-tilt
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findTilt(root *TreeNode) int {
	var (
		sumTilt int
		traverse func(root *TreeNode) int
	)
	absInt := func(num int) int {
		if num < 0 {
			return -num
		}
		return num
	}
	traverse = func(root *TreeNode) int{
		if root == nil {
			return 0
		}
		left := traverse(root.Left)
		right := traverse(root.Right)
		sumTilt += absInt(left - right)
		return left + right + root.Val
	}
	traverse(root)
	return sumTilt
}




}
