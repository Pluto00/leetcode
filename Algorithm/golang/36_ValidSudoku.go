package main

/*
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。

数独部分空格内已填入了数字，空白格用 '.' 表示。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/valid-sudoku
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func isValidSudoku(board [][]byte) bool {
	var (
		rows    [10][10]int
		columns [10][10]int
		boxes   [10][10]int
	)
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] != '.' {
				idx := board[i][j] - '0'
				rows[i][idx]++
				columns[idx][j]++
				boxes[(i/3)*3+j/3][idx]++
				if rows[i][idx] > 1 || columns[idx][j] > 1 || boxes[(i/3)*3+j/3][idx] > 1 {
					return false
				}
			}
		}
	}
	return true
}
