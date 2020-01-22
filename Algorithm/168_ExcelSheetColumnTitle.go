package main

/*
给定一个正整数，返回它在 Excel 表中相对应的列名称。

例如，

    1 -> A
    2 -> B
    3 -> C
    ...
    26 -> Z
    27 -> AA
    28 -> AB
    ...

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/excel-sheet-column-title
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func convertToTitle(n int) string {
	ans := []uint8{}
	for n > 0 {
		n--
		ans = append([]uint8{uint8('A' + n%26)}, ans...)
		n /= 26
	}
	return string(ans)
}
