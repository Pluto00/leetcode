package main

/*
给定一个整数数组和一个整数 k，判断数组中是否存在两个不同的索引 i 和 j，使得 nums [i] = nums [j]，并且 i 和 j 的差的绝对值最大为 k。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		if j, ok := m[nums[i]]; ok && i-j <= k {
			return true
		}
		m[nums[i]] = i
	}
	return false
}
