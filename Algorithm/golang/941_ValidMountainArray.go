package main

/*
给定一个整数数组 A，如果它是有效的山脉数组就返回 true，否则返回 false。

让我们回顾一下，如果 A 满足下述条件，那么它是一个山脉数组：

A.length >= 3
在 0 < i < A.length - 1 条件下，存在 i 使得：
A[0] < A[1] < ... A[i-1] < A[i]
A[i] > A[i+1] > ... > A[B.length - 1]

提示：
0 <= A.length <= 10000
0 <= A[i] <= 10000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-mountain-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func validMountainArray(A []int) bool {
	i := 0
	for ; i < len(A)-1; i++ {
		if A[i] > A[i+1] {
			break
		}
	}
	if i == 0 || i == len(A)-1 {
		return false
	}
	for i++; i < len(A); i++ {
		if A[i] >= A[i-1] {
			return false
		}
	}
	return true
}
