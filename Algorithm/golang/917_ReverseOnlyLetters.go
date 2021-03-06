package main

/*
给定一个字符串 S，返回 “反转后的” 字符串，其中不是字母的字符都保留在原地，而所有字母的位置发生反转。

示例 1：

输入："ab-cd"
输出："dc-ba"
示例 2：

输入："a-bC-dEf-ghIj"
输出："j-Ih-gfE-dCba"
示例 3：

输入："Test1ng-Leet=code-Q!"
输出："Qedo1ct-eeLg=ntse-T!"

提示：

S.length <= 100
33 <= S[i].ASCIIcode <= 122
S 中不包含 \ or "

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-only-letters
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func reverseOnlyLetters(S string) string {
	var strBytes = []byte(S)
	var l, r = 0, len(strBytes) - 1
	isLetter := func(x byte) bool {
		return ('a' <= x && x <= 'z') || ('A' <= x && x <= 'Z')
	}
	for l < r {
		for l < len(strBytes) && !isLetter(strBytes[l]) {
			l++
		}
		for r >= 0 && !isLetter(strBytes[r]) {
			r--
		}
		if l < r {
			strBytes[l], strBytes[r] = strBytes[r], strBytes[l]
		}
		l++
		r--
	}
	return string(strBytes)
}
