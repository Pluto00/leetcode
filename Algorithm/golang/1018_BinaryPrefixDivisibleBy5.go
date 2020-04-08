package main

/*
给定由若干 0 和 1 组成的数组 A。我们定义 N_i：从 A[0] 到 A[i] 的第 i 个子数组被解释为一个二进制数（从最高有效位到最低有效位）。

返回布尔值列表 answer，只有当 N_i 可以被 5 整除时，答案 answer[i] 为 true，否则为 false。

提示：

1 <= A.length <= 30000
A[i] 为 0 或 1

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-prefix-divisible-by-5
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func prefixesDivBy5(A []int) []bool {
	var ans []bool
	var sum int
	for i := 0; i < len(A); i++ {
		sum = (sum<<1 + A[i]) % 5
		ans = append(ans, sum == 0)
	}
	return ans
}
