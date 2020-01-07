package main

/*
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

说明:

必须在原数组上操作，不能拷贝额外的数组。
尽量减少操作次数。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/move-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func moveZeroes(nums []int) {
	length := len(nums)
	for i := 0; i < len(nums)-1; i++ {
		for i < len(nums)-1 && nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
		}
	}
	nums = append(nums, make([]int, length-len(nums))...)
}
