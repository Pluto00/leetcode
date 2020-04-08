package main

/*
给定一个二叉树，检查它是否是镜像对称的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/symmetric-tree
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
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return isSymmetricTree(root.Left, root.Right)
}
func isSymmetricTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p==q
	}
	if p.Val != q.Val {
		return false
	}
	return isSymmetricTree(p.Left, q.Right) && isSymmetricTree(p.Right, q.Left)
}
