package main

/*
给定一个由 '1'（陆地）和 '0'（水）组成的的二维网格，计算岛屿的数量。一个岛被水包围，并且它是通过水平方向或垂直方向上相邻的陆地连接而成的。你可以假设网格的四个边均被水包围。

示例 1:

输入:
11110
11010
11000
00000

输出: 1
示例 2:

输入:
11000
11000
00100
00011

输出: 3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/number-of-islands
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var (
		step = [4][2]int{
			[2]int{0, 1},
			[2]int{1, 0},
			[2]int{0, -1},
			[2]int{-1, 0},
		}
		n       = len(grid)
		m       = len(grid[0])
		visited = make([][]bool, n)
		part    = 0
		dfs     func(int, int)
	)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	dfs = func(r, c int) {
		if r < 0 || r >= n || c < 0 || c >= m {
			return
		}
		if grid[r][c] == '0' || visited[r][c] {
			return
		}
		visited[r][c] = true
		for i := 0; i < 4; i++ {
			dfs(r+step[i][0], c+step[i][1])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '1' && !visited[i][j] {
				part++
				dfs(i, j)
			}
		}
	}
	return part
}
