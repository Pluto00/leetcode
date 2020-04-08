package main

/*
给你一棵二叉树，请你返回层数最深的叶子节点的和。

示例：

输入：root = [1,2,3,4,5,null,6,7,null,null,null,null,8]
输出：15

提示：

树中节点数目在 1 到 10^4 之间。
每个节点的值在 1 到 100 之间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/deepest-leaves-sum
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
func deepestLeavesSum(root *TreeNode) int {
	var (
		maxDep = -1
		ans    int
		dfs    func(*TreeNode, int)
	)
	dfs = func(node *TreeNode, dep int) {
		if node == nil {
			return
		}
		if dep == maxDep {
			ans += node.Val
		} else if dep > maxDep {
			maxDep = dep
			ans = node.Val
		}
		dfs(node.Left, dep+1)
		dfs(node.Right, dep+1)
	}
	dfs(root, 0)
	return ans
}
