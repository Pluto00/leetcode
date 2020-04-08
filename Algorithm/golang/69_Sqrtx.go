package main

/*
实现 int sqrt(int x) 函数。

计算并返回 x 的平方根，其中 x 是非负整数。

由于返回类型是整数，结果只保留整数的部分，小数部分将被舍去。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sqrtx
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/


/*
牛顿迭代法：sqrt(a)
f(x) = x ^ 2 - a
f'(x) = 2 * x
x_n+1 = x_n - f(x_n) / f'(x_n)
      = (x_n + a/x_n) / 2
*/

//二分
func mySqrt(x int) int {
	var l, r = 0, x
	for i := (l + r) / 2; l < r; i = (l + r) / 2 {
		if i*i < x {
			l = i + 1
		} else {
			r = i
		}
	}
	if l*l > x {
		return l - 1
	}
	return l
}
