package main

import "strings"

/*
将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。

比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：

L   C   I   R
E T O E S I I G
E   D   H   N
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。

请你实现这个将字符串进行指定行数变换的函数：

string convert(string s, int numRows);

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/zigzag-conversion
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	var (
		goDown = false
		rows   = make([]string, numRows)
		curRow = 0
	)
	for i := 0; i < len(s); i++ {
		rows[curRow] += string(s[i])
		if curRow == 0 || curRow == numRows-1 {
			goDown = !goDown
		}
		if goDown {
			curRow++
		} else {
			curRow--
		}
	}
	return strings.Join(rows, "")
}
