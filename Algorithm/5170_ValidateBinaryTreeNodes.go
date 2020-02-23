package main

/*
5170. 验证二叉树 显示英文描述
用户通过次数0
用户尝试次数0
通过次数0
提交次数0
题目难度Medium
二叉树上有 n 个节点，按从 0 到 n - 1 编号，其中节点 i 的两个子节点分别是 leftChild[i] 和 rightChild[i]。

只有 所有 节点能够形成且 只 形成 一颗 有效的二叉树时，返回 true；否则返回 false。

如果节点 i 没有左子节点，那么 leftChild[i] 就等于 -1。右子节点也符合该规则。

注意：节点没有值，本问题中仅仅使用节点编号。

示例 1：

输入：n = 4, leftChild = [1,-1,3,-1], rightChild = [2,-1,-1,-1]
输出：true
示例 2：

输入：n = 4, leftChild = [1,-1,3,-1], rightChild = [2,3,-1,-1]
输出：false
示例 3：

输入：n = 2, leftChild = [1,0], rightChild = [-1,-1]
输出：false
示例 4：

输入：n = 6, leftChild = [1,-1,-1,4,-1,-1], rightChild = [2,-1,-1,5,-1,-1]
输出：false

提示：

1 <= n <= 10^4
leftChild.length == rightChild.length == n
-1 <= leftChild[i], rightChild[i] <= n - 1
*/

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	var (
		visited []bool = make([]bool, n)
		dfs     func(int)
		ans     bool = true
	)
	dfs = func(root int) {
		if visited[root] {
			ans = false
			return
		}
		visited[root] = true
		if leftChild[root] != -1 {
			dfs(leftChild[root])
		}
		if rightChild[root] != -1 {
			dfs(rightChild[root])
		}
	}
	dfs(0)
	for i := 0; i < n; i++ {
		if !visited[i] {
			ans = false
			break
		}
	}
	return ans
}
