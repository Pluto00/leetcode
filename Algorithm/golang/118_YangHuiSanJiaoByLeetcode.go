package main

/*
给定一个非负整数 numRows，生成杨辉三角的前 numRows 行。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/pascals-triangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func generate(numRows int) [][]int {
	if numRows == 0 {
		return nil
	} 
	var ans [][]int  // nil
	ans = make([][]int, numRows)  // ans := make([][]int, numRows) 
	for i:=0;i<numRows ; i++ {
		ans[i] = make([]int, i+1)
		ans[i][0] = 1
		for j:=1;j<i;j++ {
			ans[i][j] = ans[i-1][j-1]+ans[i-1][j]
		}
		ans[i][i] = 1
	}
	return ans
}