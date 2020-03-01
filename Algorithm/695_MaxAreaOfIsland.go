package main

/*
给定一个包含了一些 0 和 1的非空二维数组 grid , 一个 岛屿 是由四个方向 (水平或垂直) 的 1 (代表土地) 构成的组合。你可以假设二维矩阵的四个边缘都被水包围着。

找到给定的二维数组中最大的岛屿面积。(如果没有岛屿，则返回面积为0。)

示例 1:

[[0,0,1,0,0,0,0,1,0,0,0,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,1,1,0,1,0,0,0,0,0,0,0,0],
 [0,1,0,0,1,1,0,0,1,0,1,0,0],
 [0,1,0,0,1,1,0,0,1,1,1,0,0],
 [0,0,0,0,0,0,0,0,0,0,1,0,0],
 [0,0,0,0,0,0,0,1,1,1,0,0,0],
 [0,0,0,0,0,0,0,1,1,0,0,0,0]]
对于上面这个给定矩阵应返回 6。注意答案不应该是11，因为岛屿只能包含水平或垂直的四个方向的‘1’。

示例 2:

[[0,0,0,0,0,0,0,0]]
对于上面这个给定的矩阵, 返回 0。

注意: 给定的矩阵grid 的长度和宽度都不超过 50。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/max-area-of-island
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func maxAreaOfIsland(grid [][]int) int {
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
		MaxArea = 0
		TmpArea = 0
		dfs     func(int, int)
	)
	for i := 0; i < n; i++ {
		visited[i] = make([]bool, m)
	}
	dfs = func(r, c int) {
		if r < 0 || r >= n || c < 0 || c >= m {
			return
		}
		if grid[r][c] == 0 || visited[r][c] {
			return
		}
		visited[r][c] = true
		TmpArea++
		for i := 0; i < 4; i++ {
			dfs(r+step[i][0], c+step[i][1])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				TmpArea = 0
				dfs(i, j)
				if TmpArea > MaxArea {
					MaxArea = TmpArea
				}

			}
		}
	}
	return MaxArea
}
