package main

/*
给定一个Excel表格中的列名称，返回其相应的列序号。

致谢：
特别感谢 @ts 添加此问题并创建所有测试用例。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/excel-sheet-column-number
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func titleToNumber(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans = ans*26 + int(s[i]-'A'+1)
	}
	return ans
}
