package main

/*
给定范围 [m, n]，其中 0 <= m <= n <= 2147483647，返回此范围内所有数字的按位与（包含 m, n 两端点）。

示例 1:

输入: [5,7]
输出: 4
示例 2:

输入: [0,1]
输出: 0

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/bitwise-and-of-numbers-range
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// solution 1：
func rangeBitwiseAnd(m int, n int) int {
	var zeros = 0
	for n > m {
		zeros++
		n >>= 1
		m >>= 1
	}
	return m << zeros
}

// solution 2:
func rangeBitwiseAnd(m int, n int) int {
	var num = m ^ n
	num |= num >> 1
	num |= num >> 2
	num |= num >> 4
	num |= num >> 8
	num |= num >> 16
	return m & (^num)
}
