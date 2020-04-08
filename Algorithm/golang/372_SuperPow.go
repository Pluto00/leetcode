package main

/*
你的任务是计算 ab 对 1337 取模，a 是一个正整数，b 是一个非常大的正整数且会以数组形式给出。

示例 1:

输入: a = 2, b = [3]
输出: 8
示例 2:

输入: a = 2, b = [1,0]
输出: 1024

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/super-pow
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func superPow(a int, b []int) int {
	var mod = 1337
	var ans = 1
	var fastPow = func(x, k int) int {
		if x <= 1 {
			return x
		}
		ret := 0
		for ; k > 0; k >>= 1 {
			if k&1 == 1 {
				ret = x * ret % mod
			}
			x = x * x % mod
		}
		return ret
	}
	for _, e := range b {
		ans = fastPow(ans, 10) * fastPow(a, e) % mod
	}
	return ans
}
