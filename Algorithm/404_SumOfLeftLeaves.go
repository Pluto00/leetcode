package main

/*
计算给定二叉树的所有左叶子之和。

示例：

    3
   / \
  9  20
    /  \
   15   7

在这个二叉树中，有两个左叶子，分别是 9 和 15，所以返回 24

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-left-leaves
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

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var ans int
	if root.Left.Right == nil && root.Left.Left == nil {
		ans += root.Left.Val
	} else {
		ans += sumOfLeftLeaves(root.Left)
	}
	ans += sumOfLeftLeaves(root.Right)
	return ans
}
