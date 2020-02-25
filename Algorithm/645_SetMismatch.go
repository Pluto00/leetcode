package main

/*
集合 S 包含从1到 n 的整数。不幸的是，因为数据错误，导致集合里面某一个元素复制了成了集合里面的另外一个元素的值，导致集合丢失了一个整数并且有一个元素重复。

给定一个数组 nums 代表了集合 S 发生错误后的结果。你的任务是首先寻找到重复出现的整数，再找到丢失的整数，将它们以数组的形式返回。

示例 1:

输入: nums = [1,2,2,4]
输出: [2,3]
注意:

给定数组的长度范围是 [2, 10000]。
给定的数组是无序的。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/set-mismatch
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func findErrorNums(nums []int) []int {
	var (
		ans = make([]int, 2)
	)
	for i := 0; i < len(nums); i++ {
		key := nums[i]
		if key < 0 {
			key = -key
		}
		if nums[key-1] < 0 {
			ans[0] = key
		} else {
			nums[key-1] = -nums[key-1]
		}
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			ans[1] = i + 1
		}
	}
	return ans
}
