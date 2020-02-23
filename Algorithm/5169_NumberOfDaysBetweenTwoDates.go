package main

import "time"

/*
请你编写一个程序来计算两个日期之间隔了多少天。

日期以字符串形式给出，格式为 YYYY-MM-DD，如示例所示。



示例 1：

输入：date1 = "2019-06-29", date2 = "2019-06-30"
输出：1
示例 2：

输入：date1 = "2020-01-15", date2 = "2019-12-31"
输出：15


提示：

给定的日期是 1971 年到 2100 年之间的有效日期。
*/

func daysBetweenDates(date1 string, date2 string) int {
	t1, _ := time.ParseInLocation("2006-01-02", date1, time.UTC)
	t2, _ := time.ParseInLocation("2006-01-02", date2, time.UTC)
	dayTime := int64(24 * 60 * 60)
	if t1.Unix()-t2.Unix() > 0 {
		return int((t1.Unix() - t2.Unix()) / dayTime)
	}
	return int((t2.Unix() - t1.Unix()) / dayTime)
}
