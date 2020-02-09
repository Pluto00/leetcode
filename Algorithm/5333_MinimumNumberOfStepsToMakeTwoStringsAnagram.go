package main

/*
给你两个长度相等的字符串 s 和 t。每一个步骤中，你可以选择将 t 中的 任一字符 替换为 另一个字符。

返回使 t 成为 s 的字母异位词的最小步骤数。

字母异位词 指字母相同，但排列不同的字符串。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/minimum-number-of-steps-to-make-two-strings-anagram
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func minSteps(s string, t string) int {
	var (
		sWords [26]int
		tWords [26]int
		abs    func(int) int
		ans    int
	)
	abs = func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}

	for i := 0; i < len(s); i++ {
		sWords[s[i]-'a']++
		tWords[t[i]-'a']++
	}
	for i := 0; i < 26; i++ {
		ans += abs(sWords[i] - tWords[i])
	}
	return ans / 2
}
