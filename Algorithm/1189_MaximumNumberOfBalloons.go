package main

/*
给你一个字符串 text，你需要使用 text 中的字母来拼凑尽可能多的单词 "balloon"（气球）。

字符串 text 中的每个字母最多只能被使用一次。请你返回最多可以拼凑出多少个单词 "balloon"。

提示：

1 <= text.length <= 10^4
text 全部由小写英文字母组成

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximum-number-of-balloons
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxNumberOfBalloons(text string) int {
	ans, count := 0x3f3f3f3f, [5]int{}
	for i := 0; i < len(text); i++ {
		if text[i] == 'b' {
			count[0]++
		} else if text[i] == 'a' {
			count[1]++
		} else if text[i] == 'l' {
			count[2]++
		} else if text[i] == 'o' {
			count[3]++
		} else if text[i] == 'n' {
			count[4]++
		}
	}
	count[2] /= 2
	count[3] /= 2
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	for i := 0; i < len(count); i++ {
		ans = min(ans, count[i])
	}
	return ans
}
