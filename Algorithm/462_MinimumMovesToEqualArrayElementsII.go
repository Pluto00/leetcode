package main

import "sort"

/*
给定一个非空整数数组，找到使所有数组元素相等所需的最小移动数，其中每次移动可将选定的一个元素加1或减1。 您可以假设数组的长度最多为10000。

例如:

输入:
[1,2,3]

输出:
2

说明：
只有两个动作是必要的（记得每一步仅可使其中一个元素加1或减1）：

[1,2,3]  =>  [2,2,3]  =>  [2,2,2]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-moves-to-equal-array-elements-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minMoves2(nums []int) int {
	sort.Ints(nums)
	var ans int
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		ans += nums[j] - nums[i]
	}
	return ans
}
