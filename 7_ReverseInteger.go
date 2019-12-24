package main

import "fmt"

/*
给出一个 32 位的有符号整数，你需要将这个整数中每位上的数字进行反转。

注意:
假设我们的环境只能存储得下 32 位的有符号整数，则其数值范围为 [−231,  231 − 1]。请根据这个假设，如果反转后整数溢出那么就返回 0。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-integer
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func reverse(x int) int {
	var ans, flag = 0, 1
	if x < 0 {
		x = -x
		flag = -1
	}
	for x > 0 {
		ans = ans*10 + x%10
		x /= 10
	}
	if ans > 1<<31-1 || ans < - 1>>31 {
		return 0
	}
	return flag * ans
}

func main() {
	fmt.Println(reverse(120))
}
