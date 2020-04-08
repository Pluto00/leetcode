package main

/*
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

说明:

你可以假设数组不可变。
会多次调用 sumRange 方法。

来源：力扣（LeetCode）
链接：https://dev.lingkou.xyz/problems/range-sum-query-immutable
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type NumArray struct {
	nums []int
}

func Constructor(nums []int) NumArray {
	if nums == nil {
		return NumArray{nums: nil}
	}
	obj := NumArray{nums: make([]int, len(nums)+1)}
	for i := 0; i < len(nums); i++ {
		obj.nums[i+1] = nums[i] + obj.nums[i]
	}
	return obj
}

func (this *NumArray) SumRange(i int, j int) int {
	return this.nums[j] - this.nums[i]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */
