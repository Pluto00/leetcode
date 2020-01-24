package main

/*
请你帮忙给从 1 到 n 的数设计排列方案，使得所有的「质数」都应该被放在「质数索引」（索引从 1 开始）上；你需要返回可能的方案总数。

让我们一起来回顾一下「质数」：质数一定是大于 1 的，并且不能用两个小于它的正整数的乘积来表示。

由于答案可能会很大，所以请你返回答案 模 mod 10^9 + 7 之后的结果即可。

提示：

1 <= n <= 100

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/prime-arrangements
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

// A(n, n) = n!
func numPrimeArrangements(n int) int {
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101}
	var (
		cnt  = 0
		sum1 = 1
		sum2 = 1
		MOD  = int(1e9 + 7)
	)
	for i := 0; primes[i] <= n; i++ {
		cnt++
	}
	for i := 1; i <= cnt; i++ {
		sum1 = (i * sum1) % MOD
	}
	for i := 1; i <= n-cnt; i++ {
		sum2 = (i * sum2) % MOD
	}
	return (sum1 * sum2) % MOD
}
