package main

import "golang.org/x/net/html/atom"

/*
给定一个未排序的整数数组，找出其中没有出现的最小的正整数。

示例 1:

输入: [1,2,0]
输出: 3
示例 2:

输入: [3,4,-1,1]
输出: 2
示例 3:

输入: [7,8,9,11,12]
输出: 1
说明:

你的算法的时间复杂度应为O(n)，并且只能使用常数级别的空间。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/first-missing-positive
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func firstMissingPositive(nums []int) int {
	var (
		length   = len(nums)
		countOf1 = 0
	)
	for i := 0; i < length; i++ {
		if nums[i] == 1 {
			countOf1++
		} else if nums[i] <= 0 || nums[i] > length {
			nums[i] = 1
		}
	}
	if countOf1 == 0 {
		return 1
	}
	if length == 1 {
		return 2
	}
	for i := 0; i < length; i++ {
		idx := nums[i]
		if idx < 0 {
			idx = -idx
		}
		idx = idx - 1
		if nums[idx] > 0 {
			nums[idx] = -nums[idx]
		}
	}
	for i := 0; i < length; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return length + 1
}
