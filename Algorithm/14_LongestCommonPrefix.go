package main

import (
	"fmt"
)
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

说明:

所有输入只包含小写字母 a-z 。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-common-prefix
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var i, j = 0, 0
	for i = 0; i < len(strs[0]); i++ {
		for j = 1; j < len(strs); j++ {
			if i >= len(strs[j]) || strs[j][i] != strs[j-1][i] {
				return strs[0][0 : i]
			}
		}
	}
	return strs[0][0 : i]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"dos"}))
}
