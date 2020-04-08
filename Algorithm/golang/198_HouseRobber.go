package main

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/house-robber
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// dp[i] = max(dp[i-2] + nums[i], dp[i-1])
func rob(nums []int) int {
	dpI2, dpI1, dpI := 0, 0, 0
	for i := 0; i < len(nums); i++ {
		dpI = max(dpI2+nums[i], dpI1)
		dpI2, dpI1 = dpI1, dpI
	}
	return dpI
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
