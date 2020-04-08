package Algorithm

/*
给定一个二叉树，返回它的 前序 遍历。

进阶: 递归算法很简单，你可以通过迭代算法完成吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-preorder-traversal
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
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var stack []*TreeNode
	var ans []int
	stack = append(stack, root)
	for len(stack) != 0 {
		tmp := stack[len(stack)-1]
		ans = append(ans, tmp.Val)
		stack = stack[:len(stack)-1]
		if tmp.Right != nil {
			stack = append(stack, tmp.Right)
		}
		if tmp.Left != nil {
			stack = append(stack, tmp.Left)
		}
	}
	return ans
}
