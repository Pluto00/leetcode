package main

import "go-common-master/app/service/main/workflow/model"

/*
泰波那契序列 Tn 定义如下：

T0 = 0, T1 = 1, T2 = 1, 且在 n >= 0 的条件下 Tn+3 = Tn + Tn+1 + Tn+2

给你整数 n，请返回第 n 个泰波那契数 Tn 的值。

提示：

0 <= n <= 37
答案保证是一个 32 位整数，即 answer <= 2^31 - 1。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/n-th-tribonacci-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n <= 2 {
		return 1
	}
	t0, t1, t2 := 0, 1, 1
	for i := 0; i < n-2; i++ {
		t0, t1, t2 = t1, t2, t0+t1+t2

	}
	return t2
}
