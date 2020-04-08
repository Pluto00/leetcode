package main

/*
给你一个整数 n 。请你先将从 1 到 n 的所有整数按 10 进制对各个数位求和，然后把数位和相等的数字放到同一个组中。

请你统计每个组的数字数目，并返回数字数目并列最多的组有多少个。



示例 1：

输入：n = 13
输出：4
解释：总共有 9 个组，将 1 到 13 按数位求和后这些组分别是：
[1,10]，[2,11]，[3,12]，[4,13]，[5]，[6]，[7]，[8]，[9]。总共有 4 个组拥有的数字并列最多。
示例 2：

输入：n = 2
输出：2
解释：总共有 2 个大小为 1 的组 [1]，[2]。
示例 3：

输入：n = 15
输出：6
示例 4：

输入：n = 24
输出：5


提示：

1 <= n <= 10^4
*/

func countLargestGroup(n int) int {
	numOfGroup := make([]int, 40)
	sumOfDigit := func(num int) int {
		ret := 0
		for num > 0 {
			ret += num % 10
			num /= 10
		}
		return ret
	}
	for i := 1; i <= n; i++ {
		numOfGroup[sumOfDigit(i)]++
	}
	largestGroup := -1
	ans := -1
	for i := 0; i < len(numOfGroup); i++ {
		if numOfGroup[i] > largestGroup {
			largestGroup = numOfGroup[i]
			ans = 1
		} else if numOfGroup[i] == largestGroup {
			ans++
		}
	}
	return ans
}
