package main

/*
给定两个整数 n 和 k，返回 1 ... n 中所有可能的 k 个数的组合。

示例:

输入: n = 4, k = 2
输出:
[
  [2,4],
  [3,4],
  [2,3],
  [1,2],
  [1,3],
  [1,4],
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/combinations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func combine(n int, k int) [][]int {
	var (
		ans  [][]int
		last int
		dfs  func(int)
		res  = make([]int, k)
	)
	dfs = func(num int) {
		if last == k {
			tmp := make([]int, k)
			copy(tmp, res)
			ans = append(ans, tmp)
			return
		}
		for i := num + 1; i <= n && n-i >= k-last-1; i++ {
			res[last] = i
			last++
			dfs(i)
			last--
		}
	}
	dfs(1)
	return ans
}
