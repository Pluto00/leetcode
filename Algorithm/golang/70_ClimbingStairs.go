package main

/*
假设你正在爬楼梯。需要 n 阶你才能到达楼顶。

每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？

注意：给定 n 是一个正整数。


来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/climbing-stairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

/*
f(n) = f(n-1) + f(n-2)
*/

func climbStairs(n int) int {
	var f1, f2 = 0, 1
	for i := 3; i <= n; i++ {
		f1, f2 = f2, f1 + f2
	}
	return f2
}
