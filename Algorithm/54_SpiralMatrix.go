package main

/*
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func spiralOrder(matrix [][]int) []int {
	var ans []int
	if len(matrix) == 0 {
		return ans
	}
	var (
		r1, r2 = 0, len(matrix) - 1
		c1, c2 = 0, len(matrix[0]) - 1
	)
	for r1 <= r2 && c1 <= c2 {
		for c := c1; c <= c2; c++ {
			ans = append(ans, matrix[r1][c])
		}
		for r := r1 + 1; r <= r2; r++ {
			ans = append(ans, matrix[r][c2])
		}
		if r1 < r2 && c1 < c2 {
			for c := c2 - 1; c > c1; c-- {
				ans = append(ans, matrix[r2][c])
			}
			for r := r2; r > r1; r-- {
				ans = append(ans, matrix[r][c1])
			}
		}
		r1++
		r2--
		c1++
		c2--
	}
	return ans
}
