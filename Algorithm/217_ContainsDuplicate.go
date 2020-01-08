package main

/*
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/contains-duplicate
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func containsDuplicate(nums []int) bool {
	m := map[int]bool{}
	for i := 0; i < len(nums); i++ {
		if _, ok := m[nums[i]]; ok {
			return ok
		} else {
			m[nums[i]] = true
		}
	}
	return false
}
