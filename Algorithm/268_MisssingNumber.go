package main

/*
给定一个包含 0, 1, 2, ..., n 中 n 个数的序列，找出 0 .. n 中没有出现在序列中的那个数。

说明:
你的算法应具有线性时间复杂度。你能否仅使用额外常数空间来实现?

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/missing-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func missingNumber(nums []int) int {
	var ans int
	for i := 0; i < len(nums); i++ {
		ans ^= (i + 1) ^ nums[i]
	}
	return ans
}
