package main

/*
给定 n 个整数，找出平均数最大且长度为 k 的连续子数组，并输出该最大平均数。

示例 1:

输入: [1,12,-5,-6,50,3], k = 4
输出: 12.75
解释: 最大平均数 (12-5-6+50)/4 = 51/4 = 12.75


注意:

1 <= k <= n <= 30,000。
所给数据范围 [-10,000，10,000]。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-average-subarray-i
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findMaxAverage(nums []int, k int) float64 {
	var sum float64
	for i := 0; i < k; i++ {
		sum += float64(nums[i])
	}
	var ans = sum
	for i := k - 1; i < len(nums); i++ {
		sum += float64(nums[i] - nums[i-k])
		if sum > ans {
			ans = sum
		}
	}
	return ans / float64(k)
}
