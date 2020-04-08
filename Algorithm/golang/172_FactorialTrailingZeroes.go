package main

/*
给定一个整数 n，返回 n! 结果尾数中零的数量。

说明: 你算法的时间复杂度应为 O(log n) 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/factorial-trailing-zeroes
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func trailingZeroes(n int) int {
	ans := 0
	for n > 0 {
		ans += n / 5
		n /= 5
	}
	return ans
}
