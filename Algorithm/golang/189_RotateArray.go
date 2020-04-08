package main

/*
给定一个数组，将数组中的元素向右移动 k 个位置，其中 k 是非负数。

说明:

尽可能想出更多的解决方案，至少有三种不同的方法可以解决这个问题。
要求使用空间复杂度为 O(1) 的 原地 算法。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func rotate(nums []int, k int) {
	k = k % len(nums)
	rotateArray(0, len(nums)-k, nums)
	rotateArray(len(nums)-k, len(nums), nums)
	rotateArray(0, len(nums), nums)
}

func rotateArray(l, r int, nums []int) {
	for l < r {
		nums[l], nums[r-1] = nums[r-1], nums[l]
		l++
		r--
	}
}
