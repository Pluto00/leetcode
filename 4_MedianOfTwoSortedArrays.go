package main

import "fmt"

/*
给定两个大小为 m 和 n 的有序数组 nums1 和 nums2。

请你找出这两个有序数组的中位数，并且要求算法的时间复杂度为 O(log(m + n))。

你可以假设 nums1 和 nums2 不会同时为空。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/median-of-two-sorted-arrays
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		return findMedianSortedArrays(nums2, nums1) // 保证nums1的长度大于nums2
	}
	var n, m = len(nums1), len(nums2)
	var l, r = 0, n
	var maxLeft, minRight int
	for l <= r {
		i := (l + r) >> 1
		j := (n+m+1)>>1 - i
		if j != 0 && i != n && nums2[j-1] > nums1[i] {
			l = i + 1
		} else if i != 0 && j != m && nums1[i-1] > nums2[j] {
			r = i - 1
		} else {
			if i == 0 {
				maxLeft = nums2[j-1]
			} else if j == 0 {
				maxLeft = nums1[i-1]
			} else {
				maxLeft = func(a int, b int) int {
					if a > b {
						return a
					}
					return b
				}(nums1[i-1], nums2[j-1])
			}
			if (m+n)%2 == 1 {
				return float64(maxLeft)
			}

			if i == n {
				minRight = nums2[j]
			} else if j == m {
				minRight = nums1[i]
			} else {
				minRight = func(a int, b int) int {
					if a < b {
						return a
					}
					return b
				}(nums1[i], nums2[j])
			}
			return float64(maxLeft+minRight) / 2.0
		}
	}
	return 0.0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1,2}, []int{3,4}))
}
