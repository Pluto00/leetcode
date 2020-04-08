package main

/*
如果数组是单调递增或单调递减的，那么它是单调的。

如果对于所有 i <= j，A[i] <= A[j]，那么数组 A 是单调递增的。 如果对于所有 i <= j，A[i]> = A[j]，那么数组 A 是单调递减的。

当给定的数组 A 是单调数组时返回 true，否则返回 false。

提示：

1 <= A.length <= 50000
-100000 <= A[i] <= 100000

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/monotonic-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isMonotonic(A []int) bool {
	cmp := func(a, b int) int {
		if a > b {
			return 1
		} else if a == b {
			return 0
		} else {
			return -1
		}
	}
	store := 0
	for i := 0; i < len(A)-1; i++ {
		c := cmp(A[i], A[i+1])
		if c != 0 {
			if c != store && store != 0 {
				return false
			}
			store = c
		}
	}
	return true
}
