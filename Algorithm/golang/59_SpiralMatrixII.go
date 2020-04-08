package main

/*
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/spiral-matrix-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func generateMatrix(n int) [][]int {
	var (
		mat        = make([][]int, n)
		num        = 1
		l, r, t, b = 0, n - 1, 0, n - 1
	)
	for i := 0; i < n; i++ {
		mat[i] = make([]int, n)
	}
	for num <= n*n {
		for i := l; i <= r; i++ {
			mat[t][i] = num
			num++
		}
		t++
		for i := t; i <= b; i++ {
			mat[i][r] = num
			num++
		}
		r--
		for i := r; i >= l; i-- {
			mat[b][i] = num
			num++
		}
		b--
		for i := b; i >= t; i-- {
			mat[i][l] = num
			num++
		}
		l++
	}
	return mat
}
