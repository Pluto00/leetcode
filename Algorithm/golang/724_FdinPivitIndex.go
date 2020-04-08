package main

/*
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。

我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。

如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。

说明:

nums 的长度范围为 [0, 10000]。
任何一个 nums[i] 将会是一个范围在 [-1000, 1000]的整数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/find-pivot-index
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func pivotIndex(nums []int) int {
	var leftSum, totalSum int
	for i := 0; i < len(nums); i++ {
		totalSum += nums[i]
	}
	for i := 0; i < len(nums); i++ {
		if totalSum-nums[i]-leftSum == leftSum {
			return i
		}
		leftSum += nums[i]
	}
	return -1
}
