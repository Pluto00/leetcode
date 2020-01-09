package main

import (
	"strconv"
	"strings"
)

/*
给定一个整数，写一个函数来判断它是否是 3 的幂次方。

进阶：
你能不使用循环或者递归来完成本题吗？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/power-of-three
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}
	s := strconv.FormatInt(int64(n), 3)
	return s[0] == '1' && strings.Count(s, "0") == len(s)-1
}
