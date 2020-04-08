package main

/*
给你一个整数数组 arr 和两个整数 k 和 threshold 。

请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。

示例 1：

输入：arr = [2,2,2,2,5,5,5,8], k = 3, threshold = 4
输出：3
解释：子数组 [2,5,5],[5,5,5] 和 [5,5,8] 的平均值分别为 4，5 和 6 。其他长度为 3 的子数组的平均值都小于 4 （threshold 的值)。

提示：

1 <= arr.length <= 10^5
1 <= arr[i] <= 10^4
1 <= k <= arr.length
0 <= threshold <= 10^4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numOfSubarrays(arr []int, k int, threshold int) int {
	ans := 0
	num := []int{0}
	arr = append(num, arr...)
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	for i := k; i < len(arr); i++ {
		if (arr[i] - arr[i-k]) >= k*threshold {
			ans++
		}
	}
	return ans
}
