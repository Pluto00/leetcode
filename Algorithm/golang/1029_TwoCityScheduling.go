package main

import "sort"

/*
公司计划面试 2N 人。第 i 人飞往 A 市的费用为 costs[i][0]，飞往 B 市的费用为 costs[i][1]。

返回将每个人都飞到某座城市的最低费用，要求每个城市都有 N 人抵达。

提示：

1 <= costs.length <= 100
costs.length 为偶数
1 <= costs[i][0], costs[i][1] <= 1000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-city-scheduling
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoCitySchedCost(costs [][]int) int {
	ans, con := 0, []int{}
	for i := 0; i < len(costs); i++ {
		ans += costs[i][0]
		con = append(con, costs[i][1]-costs[i][0])
	}
	sort.Ints(con)
	for i := 0; i < len(con)>>1; i++ {
		ans += con[i]
	}
	return ans
}
