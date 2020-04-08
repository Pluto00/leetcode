package main

/*
给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/letter-combinations-of-a-phone-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var (
		table = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
		ans   = []string{""}
	)
	for i := 0; i < len(digits); i++ {
		var tmp []string
		t := table[digits[i]-'0']
		for j := 0; j < len(t); j++ {
			for k := 0; k < len(ans); k++ {
				tmp = append(tmp, ans[k]+string(t[j]))
			}
		}
		ans = tmp
	}
	return ans
}
