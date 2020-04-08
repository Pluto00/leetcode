package main

/*
给定两个有序整数数组 nums1 和 nums2，将 nums2 合并到 nums1 中，使得 num1 成为一个有序数组。

说明:

初始化 nums1 和 nums2 的元素数量分别为 m 和 n。
你可以假设 nums1 有足够的空间（空间大小大于或等于 m + n）来保存 nums2 中的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func merge(nums1 []int, m int, nums2 []int, n int) {
	nums1 = append(nums1[:m], nums2[:n]...)
	if n == 0 || m == 0 {
		return
	}
	var i, l, r = 0, m, m
	for i < l && r < m+n {
		for i < m && nums1[i] <= nums1[r] {
			i++
		}
		for r < m+n && nums1[i] >= nums1[r] {
			r++
		}
		convert(nums1, i, l, r-1)
		i += r - l + 1
		l = r
	}
}

/*
手摇算法，内存翻转算法，原地归并
nums[l....mid....r] -> nums[r.....mid....l]
*/
func convert(nums []int, l, mid, r int) {
	reverse(nums, l, mid-1)
	reverse(nums, mid, r)
	reverse(nums, l, r)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}
