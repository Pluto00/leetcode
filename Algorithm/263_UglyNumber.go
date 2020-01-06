package main

import "go-common-master/app/job/main/videoup/model/archive"

/*
编写一个程序判断给定的数是否为丑数。

丑数就是只包含质因数 2, 3, 5 的正整数。

说明：

1 是丑数。
输入不会超过 32 位有符号整数的范围: [−231,  231 − 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/ugly-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isUgly(num int) bool {
	if num == 1 {
		return true
	}
	if num%2 == 0 {
		return isUgly(num / 2)
	} else if num%3 == 0 {
		return isUgly(num / 3)
	} else if num%5 == 0 {
		return isUgly(num / 5)
	}
	return false
}
