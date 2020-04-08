package main

/*
给定一个非负整数 num，反复将各个位上的数字相加，直到结果为一位数。

进阶:
你可以不使用循环或者递归，且在 O(1) 时间复杂度内解决这个问题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-digits
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// 123 = 1 + 2 + 3 + 1 * 99 + 2 * 9
// 123 % 9 = 1 + 2 + 3
func addDigits(num int) int {
	return (num-1)%9 + 1
}
