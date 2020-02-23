package main

/*
给你一个整数 num，请你找出同时满足下面全部要求的两个整数：

两数乘积等于  num + 1 或 num + 2
以绝对差进行度量，两数大小最接近
你可以按任意顺序返回这两个整数。

示例 1：

输入：num = 8
输出：[3,3]
解释：对于 num + 1 = 9，最接近的两个因数是 3 & 3；对于 num + 2 = 10, 最接近的两个因数是 2 & 5，因此返回 3 & 3 。
示例 2：

输入：num = 123
输出：[5,25]
示例 3：

输入：num = 999
输出：[40,25]

提示：

1 <= num <= 10^9
*/

func closestDivisors(num int) []int {
	helper := func(num int) []int {
		var (
			ret    = []int{1, num}
			absRet = num - 1
		)
		for i := 2; i*i <= num; i++ {
			if num%i == 0 && num/i-i < absRet {
				ret[0], ret[1] = i, num/i
			}
		}
		return ret
	}
	ans1 := helper(num + 1)
	ans2 := helper(num + 2)
	if ans1[1]-ans1[0] < ans2[1]-ans2[0] {
		return ans1
	}
	return ans2
}
