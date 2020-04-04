package main

/*
给你一个以 (radius, x_center, y_center) 表示的圆和一个与坐标轴平行的矩形 (x1, y1, x2, y2)，其中 (x1, y1) 是矩形左下角的坐标，(x2, y2) 是右上角的坐标。

如果圆和矩形有重叠的部分，请你返回 True ，否则返回 False 。

换句话说，请你检测是否 存在 点 (xi, yi) ，它既在圆上也在矩形上（两者都包括点落在边界上的情况）。

示例 1：

输入：radius = 1, x_center = 0, y_center = 0, x1 = 1, y1 = -1, x2 = 3, y2 = 1
输出：true
解释：圆和矩形有公共点 (1,0)
示例 2：

输入：radius = 1, x_center = 0, y_center = 0, x1 = -1, y1 = 0, x2 = 0, y2 = 1
输出：true
示例 3：

输入：radius = 1, x_center = 1, y_center = 1, x1 = -3, y1 = -3, x2 = 3, y2 = 3
输出：true
示例 4：

输入：radius = 1, x_center = 1, y_center = 1, x1 = 1, y1 = -3, x2 = 2, y2 = -1
输出：false

提示：

1 <= radius <= 2000
-10^4 <= x_center, y_center, x1, y1, x2, y2 <= 10^4
x1 < x2
y1 < y2
*/

func checkOverlap(radius int, x_center int, y_center int, x1 int, y1 int, x2 int, y2 int) bool {
	//圆心在矩形内
	if x1 <= x_center && x_center <= x2 && y1 <= y_center && y_center <= y2 {
		return true
	}
	// 否则
	abs := func(a int) int {
		if a > 0 {
			return a
		}
		return -a
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	xDis := min(abs(x_center-x1), abs(x_center-x2))
	yDis := min(abs(y_center-y1), abs(y_center-y2))
	if y1 <= y_center && y_center <= y2 && radius >= xDis {
		return true
	} else if x1 <= x_center && x_center <= x2 && radius >= yDis {
		return true
	} else if radius*radius >= (xDis*xDis + yDis*yDis) {
		return true
	}
	return false
}
