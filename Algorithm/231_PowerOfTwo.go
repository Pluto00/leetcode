package main

/*
给定一个整数，编写一个函数来判断它是否是 2 的幂次方。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-two
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPowerOfTwo(n int) bool {
	return n>0 && n&(n-1) == 0
}
