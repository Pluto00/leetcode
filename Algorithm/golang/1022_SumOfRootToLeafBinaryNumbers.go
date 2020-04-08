package main

/*
给出一棵二叉树，其上每个结点的值都是 0 或 1 。每一条从根到叶的路径都代表一个从最高有效位开始的二进制数。例如，如果路径为 0 -> 1 -> 1 -> 0 -> 1，那么它表示二进制数 01101，也就是 13 。

对树上的每一片叶子，我们都要找出从根到该叶子的路径所表示的数字。

以 10^9 + 7 为模，返回这些数字之和。

提示：

树中的结点数介于 1 和 1000 之间。
node.val 为 0 或 1 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sum-of-root-to-leaf-binary-numbers
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
func sumRootToLeaf(root *TreeNode) int {
	var (
		ans      int
		traverse func(root *TreeNode, sum int)
		mod      int
	)
	mod = 1e9 + 7

	traverse = func(root *TreeNode, sum int) {
		if root.Left == nil && root.Right == nil {
			ans = (ans + sum) % mod
		}
		if root.Left != nil {
			traverse(root.Left, sum<<1+root.Val)
		}
		if root.Right != nil {
			traverse(root.Right, sum<<1+root.Val)
		}
	}

	traverse(root, 0)
	return ans
}
