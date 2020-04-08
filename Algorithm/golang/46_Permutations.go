package main

/*
给定一个没有重复数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/permutations
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func permute(nums []int) [][]int {
	var (
		result  [][]int
		res     = make([]int, len(nums))
		visited = make([]bool, len(nums))
		dfs     func(int)
	)
	dfs = func(depth int) {
		if depth == len(nums) {
			tmp := make([]int, len(res))
			copy(tmp, res)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				visited[i] = true
				res[depth] = nums[i]
				dfs(depth + 1)
				visited[i] = false
			}
		}
	}
	dfs(0)
	return result
}
